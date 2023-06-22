package collector

// Collector is an object which is able to collect some kind of facts about its
// host system.
type Collector interface {
	// GetData ensures the Collector has collected its facts.
	GetData() (map[string]string, error)
	// Flush ensures Collector has deleted previously collected data, if any.
	Flush()
}
