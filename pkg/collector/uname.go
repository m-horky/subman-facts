package collector

import (
	"git.sr.ht/~spc/go-log"
	"os"
)

// UnameCollector contains facts from 'uname'
type UnameCollector struct {
	data map[string]string
}

func (c *UnameCollector) String() string {
	return "uname collector"
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c *UnameCollector) Flush() {
	c.data = make(map[string]string)
}

// GetData collects 'uname' data.
func (c *UnameCollector) GetData() (map[string]string, error) {
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
func (c *UnameCollector) collect() error {
	fInfo, err := os.Stat("/usr/bin/uname")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname: %s", err)
		return nil
	}
	if fInfo.Mode()&0111 == 0 {
		log.Errorf("Could not read /usr/bin/uname: file not readable (%s)", fInfo.Mode())
		return nil
	}

	output, err := getCommandOutput("/usr/bin/uname", "--kernel-name")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --kernel-name: %s", err)
	} else {
		c.data["uname.sysname"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--machine")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --machine: %s", err)
	} else {
		c.data["uname.machine"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--nodename")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --nodename: %s", err)
	} else {
		c.data["uname.nodename"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--kernel-release")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --kernel-release: %s", err)
	} else {
		c.data["uname.release"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--kernel-version")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --kernel-version: %s", err)
	} else {
		c.data["uname.version"] = output
	}

	return nil
}
