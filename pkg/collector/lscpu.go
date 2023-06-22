package collector

type LscpuCollector struct {
	data map[string]string
}

func (c *LscpuCollector) String() string {
	return "lscpu collector"
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c *LscpuCollector) Flush() {
	c.data = make(map[string]string)
}

// GetData collects network interface data.
func (c *LscpuCollector) GetData() (map[string]string, error) {
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
func (c *LscpuCollector) collect() error {
	// TODO 'lscpu -e --json'
	return nil
}
