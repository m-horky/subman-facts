package facts

import "fmt"

type AzureFacts struct {
	CollectedFacts `json:"-"`
}

type AzureCollector struct {
	data      AzureFacts
	collected bool
}

func (c *AzureCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = AzureFacts{}
	}

	err := c.collect()
	if err != nil {
		return AzureFacts{}, err
	}
	return c.data, nil
}

func (c *AzureCollector) collect() error {
	return fmt.Errorf("Azure collection is not implemented.")
}
