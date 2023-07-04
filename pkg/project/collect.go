package project

import "github.com/m-horky/subman-facts/pkg/collector"

func CollectAll() (map[string]collector.CollectedFacts, []error) {
	collectors := map[string]collector.Collector{
		"distribution": &collector.DistributionCollector{},
		"memory":       &collector.MemoryCollector{},
		"uname":        &collector.UnameCollector{},
		"virt":         &collector.VirtCollector{},
	}

	collectedFacts := make(map[string]collector.CollectedFacts)
	var errors []error
	for collectorPrefix, factCollector := range collectors {
		data, err := factCollector.GetData(true)
		if err != nil {
			errors = append(errors, err)
		}
		collectedFacts[collectorPrefix] = data
	}

	return collectedFacts, errors
}
