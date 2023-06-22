package collector

import "fmt"

type NetworkCollector struct {
	data map[string]string
}

func NewNetworkCollector() NetworkCollector {
	return NetworkCollector{}
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c NetworkCollector) Flush() {
	c.data = nil
}

// GetData collects data from 'ip' via shelling out to its binary.
func (c NetworkCollector) GetData() (map[string]string, error) {
	if c.data == nil {
		err := c.collect()
		if err != nil {
			return nil, err
		}
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c NetworkCollector) collect() error {
	fmt.Println("collecting")
	return nil
}
