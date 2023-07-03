package collector

import (
	"git.sr.ht/~spc/go-log"
	"strings"
)

type VirtCollector struct {
	data map[string]string
}

func (c *VirtCollector) String() string {
	return "virt collector"
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c *VirtCollector) Flush() {
	c.data = make(map[string]string)
}

// GetData collects virtualization data.
func (c *VirtCollector) GetData() (map[string]string, error) {
	if c.data == nil {
		c.Flush()
	}
	if len(c.data) == 0 {
		err := c.collect()
		if err != nil {
			return nil, err
		}
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *VirtCollector) collect() error {
	err := c.collectVirtWhat()
	if err == nil {
		return nil
	}
	log.Errorf("Could not collect data from virt-what: %s", err)

	// If virt-what is not installed, do not do anything (RHBZ 768397)
	c.data["virt.is_guest"] = "Unknown"
	return nil
}

func (c *VirtCollector) collectVirtWhat() error {
	output, err := getCommandOutput("/usr/sbin/virt-what")
	if err != nil {
		return err
	}

	// Force into a single line, as xen can report xen and xen-hvm (RHBZ 1018807)
	output = strings.Join(strings.Split(output, "\n"), ", ")

	if len(output) == 0 {
		c.data["virt.is_guest"] = "False"
		c.data["virt.host_type"] = "Not Applicable"
		return nil
	}

	c.data["virt.is_guest"] = "True"
	c.data["virt.host_type"] = output

	// xen dom0 is a guest for virt-what's purposes, but is a host for our purposes (RHBZ 757697).
	if output == "dom0" {
		c.data["virt.is_guest"] = "False"
	}

	return nil
}

func (c *VirtCollector) collectUUID() error {

	return nil
}
