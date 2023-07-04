package collector

import "fmt"

type UptimeFacts struct {
	CollectedFacts `json:"-"`
}

type UptimeCollector struct {
	data      UptimeFacts
	collected bool
}

func (c *UptimeCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = UptimeFacts{}
	}

	err := c.collect()
	if err != nil {
		return UptimeFacts{}, err
	}
	return c.data, nil
}

func (c *UptimeCollector) collect() error {
	return fmt.Errorf("Uptime collection is not implemented.")
}
