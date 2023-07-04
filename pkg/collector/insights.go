package collector

import "fmt"

type InsightsFacts struct {
	CollectedFacts `json:"-"`
}

type InsightsCollector struct {
	data      InsightsFacts
	collected bool
}

func (c *InsightsCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = InsightsFacts{}
	}

	err := c.collect()
	if err != nil {
		return InsightsFacts{}, err
	}
	return c.data, nil
}

func (c *InsightsCollector) collect() error {
	return fmt.Errorf("Insights collection is not implemented.")
}
