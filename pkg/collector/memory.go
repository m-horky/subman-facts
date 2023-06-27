package collector

import (
	"regexp"
)

type MemoryCollector struct {
	data map[string]string
}

func (c *MemoryCollector) String() string {
	return "memory collector"
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c *MemoryCollector) Flush() {
	c.data = make(map[string]string)
}

// GetData collects memory data.
func (c *MemoryCollector) GetData() (map[string]string, error) {
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
func (c *MemoryCollector) collect() error {
	parser := regexp.MustCompile(`^(?P<key>\S*):\s*(?P<value>\d*)\s*kB`)

	lines, err := getFileOutput("/proc/meminfo")
	if err != nil {
		return err
	}
	for _, line := range lines {
		matches := parser.FindStringSubmatch(line)
		if len(matches) < 3 {
			continue
		}
		key := matches[1]
		value := matches[2]

		if key == "MemTotal" {
			c.data["memory.memtotal"] = value
		}
		// if key == "SwapTotal" {
		// 	c.data["memory.swaptotal"] = value
		// }
	}
	return nil
}
