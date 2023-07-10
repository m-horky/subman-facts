package facts

import "git.sr.ht/~spc/go-log"

type SystemFacts struct {
	CertificateVersion string `json:"certificate_version"`
	DefaultLocale      string `json:"default_locale,omitempty"`
}

// SystemCollector contains facts from 'System'
type SystemCollector struct {
	data      SystemFacts
	collected bool
}

// GetData collects certificate and locale data and returns them as SystemFacts.
func (c *SystemCollector) GetData(rescan bool) (SystemFacts, error) {
	if rescan || !c.collected {
		c.data = SystemFacts{}
	}

	err := c.collect()
	if err != nil {
		return SystemFacts{}, err
	}
	return c.data, nil
}

// collect starts the actual fact collection. It is usually invoked by GetData.
func (c *SystemCollector) collect() error {
	_ = c.collectCertificateVersion()
	_ = c.collectDefaultLocale()
	c.collected = true
	return nil
}

func (c *SystemCollector) collectCertificateVersion() error {
	// FIXME subscription-manager hardcodes this value to be '3.2'.
	// TODO Find out which certificate it should be (/etc/pki/entitlement/ or
	//  /etc/pki/consumer/)
	// see src/rhsmlib/facts/collector.py::StaticFactsCollector
	c.data.CertificateVersion = "3.2"
	log.Warn("System's certificate version fact is hardcoded, consider loading it dynamically.")
	return nil
}

func (c *SystemCollector) collectDefaultLocale() error {
	// see src/rhsmlib/facts/host_collector.py
	log.Debug("Locale fact collection is not implemented.")
	return nil
}
