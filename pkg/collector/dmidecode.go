package collector

type DmidecodeCollector struct {
	data map[string]string
}

func NewDmidecodeCollector() DmidecodeCollector {
	return DmidecodeCollector{}
}

func (c DmidecodeCollector) String() string {
	return "Dmidecode collector"
}

// Flush ensures Collector has deleted previously collected data, if any.
func (c DmidecodeCollector) Flush() {
	c.data = nil
}

// GetData collects data from 'dmidecode' via shelling out to its binary.
func (c DmidecodeCollector) GetData() (map[string]string, error) {
	if c.data == nil {
		err := c.collect()
		if err != nil {
			return nil, err
		}
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c DmidecodeCollector) collect() error {
	return nil
}
