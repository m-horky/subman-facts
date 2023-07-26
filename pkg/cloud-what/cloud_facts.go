package cloud_what

import "github.com/m-horky/subman-facts/pkg/facts"

// systemIsVM detects the system using facts.VirtCollector.
func systemIsVM() bool {
	collector := facts.NewVirtCollector()
	virtFacts, err := collector.GetData(false)
	if err != nil {
		// FIXME What should be the result here?
		return false
	}
	return virtFacts.IsGuest
}
