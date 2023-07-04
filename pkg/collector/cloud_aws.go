package collector

import "fmt"

type AWSFacts struct {
	CollectedFacts `json:"-"`
}

type AWSCollector struct {
	data      AWSFacts
	collected bool
}

func (c *AWSCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = AWSFacts{}
	}

	err := c.collect()
	if err != nil {
		return AWSFacts{}, err
	}
	return c.data, nil
}

func (c *AWSCollector) collect() error {
	return fmt.Errorf("AWS collection is not implemented.")
}
