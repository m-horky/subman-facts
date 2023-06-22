package main

import (
	"fmt"
	"github.com/m-horky/subman-facts/pkg/collector"
	"github.com/m-horky/subman-facts/pkg/project"
)

func main() {
	fmt.Printf("subman-facts, version %s\n", project.Version)
	collectors := []collector.Collector{
		collector.NewNetworkCollector(),
		collector.NewDmidecodeCollector(),
	}
	var errors []error
	for _, factCollector := range collectors {
		data, err := factCollector.GetData()
		if err != nil {
			errors = append(errors, err)
		}
		for k, v := range data {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
	for _, err := range errors {
		fmt.Println(err)
	}
}
