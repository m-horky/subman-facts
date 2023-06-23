package main

import (
	"fmt"
	"github.com/m-horky/subman-facts/pkg/collector"
	"github.com/m-horky/subman-facts/pkg/project"
	"sort"
)

func main() {
	fmt.Printf("subman-facts, version %s\n", project.Version)

	collectors := []collector.Collector{
		&collector.NetworkCollector{},
		&collector.DmidecodeCollector{},
		&collector.LscpuCollector{},
		&collector.DistributionCollector{},
	}
	var errors []error
	for _, factCollector := range collectors {
		data, err := factCollector.GetData()
		if err != nil {
			errors = append(errors, err)
		}

		// Ensure the collector's keys are sorted
		keys := make([]string, 0, len(data))
		for k, _ := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("%s: %s\n", k, data[k])
		}
	}
	for _, err := range errors {
		fmt.Println(err)
	}
}
