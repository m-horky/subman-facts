package facts

import (
	"git.sr.ht/~spc/go-log"
	"regexp"
	"strconv"
)

type MemoryFacts struct {
	CollectedFacts `json:"-"`
	MemTotal       int `json:"memtotal"`
	SwapTotal      int `json:"swaptotal"`
}

type MemoryCollector struct {
	data      MemoryFacts
	collected bool
}

// GetData collects memory data.
func (c *MemoryCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = MemoryFacts{}
	}

	err := c.collect()
	if err != nil {
		return nil, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *MemoryCollector) collect() error {
	parser := regexp.MustCompile(`^(?P<key>\S*):\s*(?P<value>\d*)\s*kB`)

	lines, err := getFileOutput("/proc/meminfo")
	if err != nil {
		log.Errorf("Could not get output of /proc/meminfo: %s")
		return err
	}
	for _, line := range lines {
		matches := parser.FindStringSubmatch(line)
		if len(matches) < 3 {
			continue
		}

		key := matches[1]
		value, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Errorf("Could not convert memory value: %s", line)
		}

		if key == "MemTotal" {
			c.data.MemTotal = value
		}
		if key == "SwapTotal" {
			c.data.SwapTotal = value
		}
	}
	return nil
}
