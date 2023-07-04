package collector

import (
	"bufio"
	"git.sr.ht/~spc/go-log"
	"os"
	"regexp"
	"strings"
)

type DistributionFacts struct {
	CollectedFacts `json:"-"`
	Name           string `json:"name"`
	Version        string `json:"version"`
	ID             string `json:"id"`
}

// DistributionCollector contains facts from /etc/os-release and /etc/redhat-release
type DistributionCollector struct {
	data      DistributionFacts
	collected bool
}

// GetData collects network interface data.
func (c *DistributionCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = DistributionFacts{}
	}

	err := c.collect()
	if err != nil {
		return nil, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *DistributionCollector) collect() error {
	_, err := os.Stat("/etc/os-release")
	if err == nil {
		err = c.collectOsRelease()
		if err == nil {
			return nil
		}
	}

	_, err = os.Stat("/etc/redhat-release")
	if err == nil {
		err = c.collectRedhatRelease()
		if err == nil {
			return nil
		}
	}
	return nil
}

// collectOsRelease collects values from /etc/os-release
func (c *DistributionCollector) collectOsRelease() error {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		log.Errorf("Could not open /etc/os-release: %s", err)
		return err
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		kv := strings.SplitN(line, "=", 2)
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

	if err := scanner.Err(); err != nil {
		log.Errorf("Could not parse /etc/os-release: %s", err)
		return err
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
