package collector

import "fmt"

type CustomFacts struct {
	CollectedFacts `json:"-"`
}

type CustomCollector struct {
	data      CustomFacts
	collected bool
}

func (c *CustomCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = CustomFacts{}
	}

	err := c.collect()
	if err != nil {
		return CustomFacts{}, err
	}
	return c.data, nil
}

func (c *CustomCollector) collect() error {
	return fmt.Errorf("Custom fact collection is not implemented.")
}
