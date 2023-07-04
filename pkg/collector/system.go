package collector

type SystemFacts struct {
	CollectedFacts     `json:"-"`
	CertificateVersion string `json:"certificate_version"`
	DefaultLocale      string `json:"default_locale"`
}

// SystemCollector contains facts from 'System'
type SystemCollector struct {
	data      SystemFacts
	collected bool
}

// GetData collects certificate and locale data and returns them as SystemFacts.
func (c *SystemCollector) GetData(rescan bool) (CollectedFacts, error) {
	if rescan || !c.collected {
		c.data = SystemFacts{}
	}

	err := c.collect()
	if err != nil {
		return nil, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *SystemCollector) collect() error {
	return nil
}

func (c *SystemCollector) collectCertificateVersion() error {
	// FIXME subscription-manager hardcodes this value to be '3.2'.
	// TODO Find out which certificate it should be (/etc/pki/entitlement/ or
	//  /etc/pki/consumer/)
	// see src/rhsmlib/facts/collector.py::StaticFactsCollector
	return nil
}

func (c *SystemCollector) collectDefaultLocale() error {
	// see src/rhsmlib/facts/host_collector.py
	return nil
}
