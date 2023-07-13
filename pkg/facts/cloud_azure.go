package facts

import "fmt"

type AzureFacts struct {
}

type AzureCollector struct {
	data      AzureFacts
	collected bool
}

func NewAzureCollector() AzureCollector {
	return AzureCollector{}
}

func (c *AzureCollector) GetData(rescan bool) (AzureFacts, error) {
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
