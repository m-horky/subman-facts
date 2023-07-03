package main

import (
	"fmt"
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/collector"
	"github.com/m-horky/subman-facts/pkg/project"
	"sort"
)

func configureLogging() {
	log.SetFlags(0)
	level, err := log.ParseLevel("debug")
	if err != nil {
		level = log.LevelError
	}
	log.SetLevel(level)
}

func main() {
	configureLogging()

	fmt.Printf("subman-facts, version %s\n", project.Version)

	collectors := []collector.Collector{
		&collector.NetworkCollector{},
		&collector.DmidecodeCollector{},
		&collector.LscpuCollector{},
		&collector.DistributionCollector{},
		&collector.UnameCollector{},
		&collector.MemoryCollector{},
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
			fmt.Print(
				fmt.Sprintf(
					"%s%s%s: ",
					"\u001B[36m", k, "\u001B[0m",
				) + fmt.Sprintf(
					"%s%s%s\n",
					"\u001B[3m", data[k], "\u001B[0m",
				),
			)
		}
	}
	for _, err := range errors {
		fmt.Println(err)
	}
}
