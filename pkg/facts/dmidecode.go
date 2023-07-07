package facts

import "fmt"

type DmidecodeFacts struct {
}

type DmidecodeCollector struct {
	data      DmidecodeFacts
	collected bool
}

func (c *DmidecodeCollector) GetData(rescan bool) (DmidecodeFacts, error) {
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
