package collector

import "fmt"

type LscpuFacts struct {
	CollectedFacts `json:"-"`
}

type LscpuCollector struct {
	data      LscpuFacts
	collected bool
}

func (c *LscpuCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = LscpuFacts{}
	}

	err := c.collect()
	if err != nil {
		return LscpuFacts{}, err
	}
	return c.data, nil
}

func (c *LscpuCollector) collect() error {
	return fmt.Errorf("Lscpu collection is not implemented.")
}
