package facts

import (
	"fmt"
	"git.sr.ht/~spc/go-log"
	"os"
	"strings"
)

type VirtFacts struct {
	IsGuest  bool   `json:"is_guest"`
	HostType string `json:"host_type"`
	UUID     string `json:"uuid,omitempty"`
}

type VirtCollector struct {
	data      VirtFacts
	collected bool
}

// GetData collects virtualization data and returns them as VirtFacts.
func (c *VirtCollector) GetData(rescan bool) (VirtFacts, error) {
	if rescan || !c.collected {
		c.data = VirtFacts{}
	}

	err := c.collect()
	if err != nil {
		return VirtFacts{}, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *VirtCollector) collect() error {
	err := c.collectVirtWhat()
	if err != nil {
		// If virt-what is not installed, do not do anything (RHBZ 768397)
		c.data.IsGuest = false
	}

	if c.data.IsGuest == true {
		_ = c.collectUUID()
	}
	c.collected = true
	return nil
}

// collectVirtWhat collects virtualization using /usr/sbin/virt-what.
func (c *VirtCollector) collectVirtWhat() error {
	output, err := getCommandOutput("/usr/sbin/virt-what")
	if err != nil {
		log.Errorf("Could not collect data from /usr/sbin/virt-what: %s", err)
		return err
	}

	// Force into a single line, as xen can report xen and xen-hvm (RHBZ 1018807)
	hosts := strings.Join(output, ", ")

	if len(hosts) == 0 {
		c.data.IsGuest = false
		c.data.HostType = "Not Applicable"
		return nil
	}

	c.data.IsGuest = true
	c.data.HostType = hosts

	// xen dom0 is a guest for virt-what's purposes, but a host for our purposes (RHBZ 757697)
	if strings.Contains(hosts, "dom0") {
		c.data.IsGuest = false
	}

	return nil
}

// collectUUID collects system UUID.
// Please note that this requires collectVirtWhat already finished to work correctly.
func (c *VirtCollector) collectUUID() error {
	// Some systems should not get their UUIDs collected (RHBZ 1438085)
	for _, hypervisor := range []string{"powervm_lx86", "xen-dom0", "ibm_systemz"} {
		if strings.Contains(c.data.HostType, hypervisor) {
			log.Debugf("We don't collect UUIDs for hypervisor '%s'.", hypervisor)
			return nil
		}
	}

	// Following implementations can overwrite the fact. For logging reasons we collect
	// the errors here and report them at the end as a whole.
	var errors []string

	// For most systems, UUID is collected by dmidecode
	err := c.collectUUIDWithDmidecode()
	if err != nil {
		errors = append(errors, err.Error())
	}

	// ppc64/ppc64le contain UUID in device-tree
	err = c.collectUUIDFromDeviceTree()
	if err != nil {
		errors = append(errors, err.Error())
	}

	// UUID may be specified on the filesystem at /sys/hypervisor/uuid
	err = c.collectUUIDFromSys()
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) == 0 {
		return nil
	}
	return fmt.Errorf("could not collect UUID from the system: %s", strings.Join(errors, ", "))
}

// collectUUIDWithDmidecode invokes the `dmidecode` binary and parses out the UUID value
func (c *VirtCollector) collectUUIDWithDmidecode() error {
	output, err := getCommandOutput("/usr/sbin/dmidecode")
	if err != nil {
		log.Errorf("Could not collect data from /usr/sbin/dmidecode: %s", err)
		return err
	}

	for _, line := range output {
		if !strings.Contains(line, "UUID") {
			continue
		}
		// The line has a format of '\tUUID: the-uuid-string'
		uuid := strings.TrimLeft(line, "\tUUID: ")
		c.data.UUID = uuid
		log.Debug("UUID value found using dmidecode.")
	}
	return nil
}

// collectUUIDFromDeviceTree collects the system UUID from device-tree platforms,
// such as ppc64 and ppc64le.
func (c *VirtCollector) collectUUIDFromDeviceTree() error {
	for _, path := range []string{"/proc/device-tree/vm,uuid", "/proc/device-tree/ibm,partition-uuid"} {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		output, err := getFileOutput(path)
		if err != nil {
			log.Errorf("UUID file %s exists, but could not be read: %s", path, err)
			return err
		}
		if len(output) < 1 {
			log.Errorf("UUID file %s does not contain readable UUID", path)
			return fmt.Errorf("could not read UUID from %s", path)
		}

		// ppc54 can report UUID with null byte at the end (RHBZ 1405125).
		uuid := strings.Trim(output[0], "\u0000")
		c.data.UUID = uuid
		log.Debugf("UUID value found in %s.", path)
	}
	return nil
}

// collectUUIDFromSys collects the system UUID from a filesystem file
// `/sys/hypervisor/uuid` (used in Xen).
func (c *VirtCollector) collectUUIDFromSys() error {
	path := "/sys/hypervisor/uuid"
	output, err := getFileOutput(path)
	if os.IsNotExist(err) {
		return nil
	}
	if len(output) < 1 {
		log.Errorf("UUID file %s does not contain readable UUID", path)
		return nil
	}

	uuid := strings.Trim(output[0], "\n\r")
	c.data.UUID = uuid
	log.Debugf("UUID value found in %s.", path)
	return nil
}
