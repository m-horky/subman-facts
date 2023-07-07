package facts

import "fmt"

type KpatchFacts struct {
}

type KpatchCollector struct {
	data      KpatchFacts
	collected bool
}

func (c *KpatchCollector) GetData(rescan bool) (KpatchFacts, error) {
	if rescan || !c.collected {
		c.data = KpatchFacts{}
	}

	err := c.collect()
	if err != nil {
		return KpatchFacts{}, err
	}
	return c.data, nil
}

func (c *KpatchCollector) collect() error {
	return fmt.Errorf("Kpatch collection is not implemented.")
}
