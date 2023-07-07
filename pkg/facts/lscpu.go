package facts

import "fmt"

type LscpuFacts struct {
}

type LscpuCollector struct {
	data      LscpuFacts
	collected bool
}

func (c *LscpuCollector) GetData(rescan bool) (LscpuFacts, error) {
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
