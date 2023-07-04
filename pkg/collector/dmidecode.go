package collector

import "fmt"

type DmidecodeFacts struct {
	CollectedFacts `json:"-"`
}

type DmidecodeCollector struct {
	data      DmidecodeFacts
	collected bool
}

func (c *DmidecodeCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = DmidecodeFacts{}
	}

	err := c.collect()
	if err != nil {
		return DmidecodeFacts{}, err
	}
	return c.data, nil
}

func (c *DmidecodeCollector) collect() error {
	return fmt.Errorf("Dmidecode collection is not implemented.")
}
