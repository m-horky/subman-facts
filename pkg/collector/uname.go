package collector

import (
	"os"
	"os/exec"
	"strings"
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
	if err != nil || fInfo.Mode()&0111 == 0 {
		// TODO Log the error or insufficient permissions
		return nil
	}

	output, err := getCommandOutput("/usr/bin/uname", "--kernel-name")
	if err != nil {
		// TODO Log the error
	} else {
		c.data["uname.sysname"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--machine")
	if err != nil {
		// TODO Log the error
	} else {
		c.data["uname.machine"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--nodename")
	if err != nil {
		// TODO Log the error
	} else {
		c.data["uname.nodename"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--kernel-release")
	if err != nil {
		// TODO Log the error
	} else {
		c.data["uname.release"] = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--kernel-version")
	if err != nil {
		// TODO Log the error
	} else {
		c.data["uname.version"] = output
	}

	return nil
}

func getCommandOutput(cmd string, args ...string) (string, error) {
	rawOutput, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(rawOutput), "\n"), nil
}
