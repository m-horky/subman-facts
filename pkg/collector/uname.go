package collector

import (
	"git.sr.ht/~spc/go-log"
	"os"
)

type UnameFacts struct {
	CollectedFacts `json:"-"`
	Sysname        string `json:"sysname"`
	Machine        string `json:"machine"`
	Nodename       string `json:"nodename"`
	KernelRelease  string `json:"release"`
	KernelVersion  string `json:"version"`
}

// UnameCollector contains facts from 'uname'
type UnameCollector struct {
	data      UnameFacts
	collected bool
}

// GetData collects 'uname' data and returns them as UnameFacts.
func (c *UnameCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = UnameFacts{}
	}

	err := c.collect()
	if err != nil {
		return nil, err
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
		c.data.Sysname = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--machine")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --machine: %s", err)
	} else {
		c.data.Machine = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--nodename")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --nodename: %s", err)
	} else {
		c.data.Nodename = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--kernel-release")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --kernel-release: %s", err)
	} else {
		c.data.KernelRelease = output
	}

	output, err = getCommandOutput("/usr/bin/uname", "--kernel-version")
	if err != nil {
		log.Errorf("Could not read /usr/bin/uname --kernel-version: %s", err)
	} else {
		c.data.KernelVersion = output
	}

	return nil
}
