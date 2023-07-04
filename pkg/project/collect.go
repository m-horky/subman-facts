package project

import (
	"fmt"
	"github.com/m-horky/subman-facts/pkg/collector"
)

func CollectAll() (map[string]collector.CollectedFacts, []error) {
	collectors := map[string]collector.Collector{
		"aws":          &collector.AWSCollector{},
		"azure":        &collector.AzureCollector{},
		"gcp":          &collector.GCPCollector{},
		"custom":       &collector.CustomCollector{},
		"distribution": &collector.DistributionCollector{},
		"dmidecode":    &collector.DmidecodeCollector{},
		"insights":     &collector.InsightsCollector{},
		"kpatch":       &collector.KpatchCollector{},
		"lscpu":        &collector.LscpuCollector{},
		"memory":       &collector.MemoryCollector{},
		"uname":        &collector.UnameCollector{},
		"uptime":       &collector.UptimeCollector{},
		"virt":         &collector.VirtCollector{},
	}

	collectedFacts := make(map[string]collector.CollectedFacts)
	var errors []error
	for collectorPrefix, factCollector := range collectors {
		data, err := factCollector.GetData(true)
		if err != nil {
			errors = append(errors, fmt.Errorf("error collecting %s: %s", collectorPrefix, err))
		}
		collectedFacts[collectorPrefix] = data
	}

	return collectedFacts, errors
}
