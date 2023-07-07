package facts

import "fmt"

type GCPFacts struct {
	CollectedFacts `json:"-"`
}

type GCPCollector struct {
	data      GCPFacts
	collected bool
}

func (c *GCPCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = GCPFacts{}
	}

	err := c.collect()
	if err != nil {
		return GCPFacts{}, err
	}
	return c.data, nil
}

func (c *GCPCollector) collect() error {
	return fmt.Errorf("GCP collection is not implemented.")
}
