package facts

import (
	"git.sr.ht/~spc/go-log"
)

type UnameFacts struct {
	Sysname       string `json:"sysname"`
	Machine       string `json:"machine"`
	Nodename      string `json:"nodename"`
	KernelRelease string `json:"release"`
	KernelVersion string `json:"version"`
}

// UnameCollector contains facts from 'uname'
type UnameCollector struct {
	data      UnameFacts
	collected bool
}

func NewUnameCollector() UnameCollector {
	return UnameCollector{
		data:      UnameFacts{},
		collected: false,
	}
}

// GetData collects 'uname' data and returns them as UnameFacts.
func (c *UnameCollector) GetData(rescan bool) (UnameFacts, error) {
	if rescan || !c.collected {
		c.data = UnameFacts{}
	}

	err := c.collect()
	if err != nil {
		return UnameFacts{}, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *UnameCollector) collect() error {
	stdout, _, err := OS.Run("/usr/bin/uname", "--kernel-name")
	if err != nil {
		log.Errorf("could not read /usr/bin/uname --kernel-name: %s", err)
	} else {
		c.data.Sysname = stdout[0]
	}

	stdout, _, err = OS.Run("/usr/bin/uname", "--machine")
	if err != nil {
		log.Errorf("could not read /usr/bin/uname --machine: %s", err)
	} else {
		c.data.Machine = stdout[0]
	}

	stdout, _, err = OS.Run("/usr/bin/uname", "--nodename")
	if err != nil {
		log.Errorf("could not read /usr/bin/uname --nodename: %s", err)
	} else {
		c.data.Nodename = stdout[0]
	}

	stdout, _, err = OS.Run("/usr/bin/uname", "--kernel-release")
	if err != nil {
		log.Errorf("could not read /usr/bin/uname --kernel-release: %s", err)
	} else {
		c.data.KernelRelease = stdout[0]
	}

	stdout, _, err = OS.Run("/usr/bin/uname", "--kernel-version")
	if err != nil {
		log.Errorf("could not read /usr/bin/uname --kernel-version: %s", err)
	} else {
		c.data.KernelVersion = stdout[0]
	}

	c.collected = true
	return nil
}
