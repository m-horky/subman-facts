package facts

import (
	"fmt"
)

type AWSFacts struct {
	InstanceID              string `json:"instance_id"`
	AccountID               string `json:"account_id"`
	BillingProducts         string `json:"billing_products"`
	MarketplaceProductCodes string `json:"marketplace_product_codes"`
}

type AWSCollector struct {
	data      AWSFacts
	collected bool
}

func (c *AWSCollector) GetData(rescan bool) (AWSFacts, error) {
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
