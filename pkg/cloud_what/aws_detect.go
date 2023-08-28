package cloud_what

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/facts"
	"github.com/nqd/flat"
	"strings"
)

// SystemRunsOn employs a strong detection of host system.
// It is the equivalent of cloud-what's `is_running_on_cloud()`.
func (_ AWSInstance) SystemRunsOn() bool {
	if systemIsVM() {
		return false
	}

	dmiCollector := facts.NewDmidecodeCollector()
	dmiFacts, err := dmiCollector.GetData(false)
	if err != nil {
		log.Info("could not get dmidecode facts, playing it safe: not running on AWS")
		return false
	}
	if strings.Contains(dmiFacts.BIOS.Version, "amazon") {
		return true
	}
	if strings.Contains(dmiFacts.BIOS.Vendor, "Amazon EC2") {
		return true
	}

	virtCollector := facts.NewVirtCollector()
	virtFacts, err := virtCollector.GetData(false)
	if err != nil {
		log.Info("could not get virt facts, playing it safe: not running on AWS")
		return false
	}
	if strings.Contains(virtFacts.HostType, "aws") {
		return true
	}

	return false
}

// SystemMaybeRunsOn employs a strong detection of host system.
// It is the equivalent of cloud-what's `is_likely_running_on_cloud()`.
// Returns the probability of the system running on this cloud.
func (i AWSInstance) SystemMaybeRunsOn() float64 {
	if systemIsVM() {
		return 0.0
	}

	virtCollector := facts.NewVirtCollector()
	virtFacts, err := virtCollector.GetData(false)
	if err != nil {
		log.Info("could not get virt facts, playing it safe: not running on AWS")
		return 0.0
	}
	dmiCollector := facts.NewDmidecodeCollector()
	dmiFacts, err := dmiCollector.GetData(false)
	if err != nil {
		log.Info("could not get dmidecode facts, playing it safe: not running on AWS")
		return 0.0
	}

	var probability = 0.0

	// We know AWS mostly uses KVM and in some cases Xen
	if strings.Contains(virtFacts.HostType, "kvm") {
		probability += 0.3
	} else if strings.Contains(virtFacts.HostType, "xen") {
		probability += 0.2
	}

	// Every UUID of a system running on AWS EC2 starts with the EC2 string
	if strings.HasPrefix(strings.ToLower(dmiFacts.System.UUID), "ec2") {
		probability += 0.1
	}

	// Strings "Amazon EC2", "Amazon" and "AWS" may be keywords in output of dmidecode
	var foundAmazonEC2 bool
	var foundAmazon bool
	var foundAWS bool
	flatennedDMIFacts, err := i.flattenDmiFacts(dmiFacts)
	if err == nil {
		for _, rawValue := range flatennedDMIFacts {
			value := strings.ToLower(fmt.Sprintf("%v", rawValue))
			if strings.Contains(value, "amazon ec2") {
				foundAmazonEC2 = true
			} else if strings.Contains(value, "amazon") {
				foundAmazon = true
			} else if strings.Contains(value, "aws") {
				foundAWS = true
			}
		}
	}
	if foundAmazonEC2 {
		probability += 0.3
	}
	if foundAmazon {
		probability += 0.2
	}
	if foundAWS {
		probability += 0.1
	}

	return probability
}

func (_ AWSInstance) flattenDmiFacts(facts facts.DmidecodeFacts) (map[string]any, error) {
	// Convert the nested structure into generic nested map
	temp, err := json.Marshal(&facts)
	if err != nil {
		log.Errorf("could not marshal DMI facts: %s", err)
		return map[string]any{}, err
	}
	var converted map[string]any
	err = json.Unmarshal(temp, &converted)
	if err != nil {
		log.Errorf("could not unmarshal marshalled DMI facts: %s", err)
		return map[string]any{}, err
	}

	flattened, err := flat.Flatten(converted, nil)
	if err != nil {
		log.Errorf("could not flatten DMI facts: %s", err)
		return map[string]any{}, err
	}
	return flattened, err
}
