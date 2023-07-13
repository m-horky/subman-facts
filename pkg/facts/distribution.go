package facts

import (
	"git.sr.ht/~spc/go-log"
	"os"
	"regexp"
	"strings"
)

type DistributionFacts struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	ID      string `json:"id"`
}

// DistributionCollector contains facts from /etc/os-release and /etc/redhat-release
type DistributionCollector struct {
	data          DistributionFacts
	collected     bool
	getFileOutput func(string) ([]string, error)
}

func NewDistributionCollector() DistributionCollector {
	return DistributionCollector{
		data:      DistributionFacts{},
		collected: false,
		getFileOutput: func(path string) ([]string, error) {
			return getFileOutput(path)
		},
	}
}

// GetData collects network interface data.
func (c *DistributionCollector) GetData(rescan bool) (DistributionFacts, error) {
	if rescan || !c.collected {
		c.data = DistributionFacts{}
	}

	err := c.collect()
	if err != nil {
		return DistributionFacts{}, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *DistributionCollector) collect() error {
	_, err := os.Stat("/etc/os-release")
	if err == nil {
		err = c.collectOsRelease()
		if err == nil {
			c.collected = true
			return nil
		}
	}

	_, err = os.Stat("/etc/redhat-release")
	if err == nil {
		err = c.collectRedhatRelease()
		if err == nil {
			c.collected = true
			return nil
		}
	}
	return nil
}

// collectOsRelease collects values from /etc/os-release
func (c *DistributionCollector) collectOsRelease() error {
	lines, err := c.getFileOutput("/etc/os-release")
	if err != nil {
		log.Errorf("Could not get output of /etc/os-release: %s", err)
		return err
	}

	for _, line := range lines {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key := kv[0]
		value := strings.Trim(kv[1], "\"")

		// 'Red Hat Enterprise Linux', 'CentOS Stream', 'Fedora Linux'
		if key == "NAME" {
			c.data.Name = value
		}
		// '8', '9', '37'
		if key == "VERSION_ID" {
			c.data.Version = value
		}
		// '9.2 (Plow)', '8.8 (Ootpa)', '9' (CentOS), '37 (Workstation Edition)'
		if key == "VERSION" {
			_ = c.collectDistributionId(value)
		}
		// 'cpe:/o:redhat:enterprise_linux:9::baseos', 'cpe:/o:centos:centos:9', 'cpe:/o:fedoraproject:fedora:37'
		if key == "CPE_NAME" {
			// TODO Reimplement Python implementation
			// vers_mod_data: List[str] = re.split(r"(?<!\\):", data["CPE_NAME"])
			// if len(vers_mod_data) >= 6:
			//     version_modifier = vers_mod_data[5].lower().replace("\\:", ":")
			log.Warn("CPE_NAME collection has not yet been implemented")
		}
	}
	return nil
}

// collectDistributionId extract proper value from VERSION field of /etc/os-release
func (c *DistributionCollector) collectDistributionId(value string) error {
	// Use only the contents of parenthesis, if possible
	re := regexp.MustCompile(`\((?P<nickname>.*?)\)`)
	matches := re.FindStringSubmatch(value)
	if len(matches) == 2 {
		value = matches[1]
	}

	c.data.ID = value
	return nil
}

// collectRedhatRelease collects value from /etc/redhat-release
func (c *DistributionCollector) collectRedhatRelease() error {
	// TODO
	return nil
}
