package facts

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CollectedFacts interface {
	IsEmpty() bool
}

// Collector is an object which is able to collect some kind of facts about its
// host system.
type Collector interface {
	// GetData ensures the Collector has collected its
	GetData(rescan bool) (CollectedFacts, error)
}

func CollectAll() (map[string]CollectedFacts, []error) {
	collectors := map[string]Collector{
		"aws":          &AWSCollector{},
		"azure":        &AzureCollector{},
		"gcp":          &GCPCollector{},
		"custom":       &CustomCollector{},
		"distribution": &DistributionCollector{},
		"dmidecode":    &DmidecodeCollector{},
		"insights":     &InsightsCollector{},
		"kpatch":       &KpatchCollector{},
		"lscpu":        &LscpuCollector{},
		"memory":       &MemoryCollector{},
		"system":       &SystemCollector{},
		"uname":        &UnameCollector{},
		"uptime":       &UptimeCollector{},
		"virt":         &VirtCollector{},
	}

	collectedFacts := make(map[string]CollectedFacts)
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

// getCommandOutput invokes a shell program and returns the output as lines.
func getCommandOutput(cmd string, args ...string) ([]string, error) {
	rawOutput, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return []string{}, err
	}
	result := strings.Split(string(rawOutput), "\n")
	return result, nil
}

// getFileOutput reads a file and returns the lines.
func getFileOutput(path string) ([]string, error) {
	rawOutput, err := os.ReadFile(path)
	if err != nil {
		return make([]string, 0), err
	}
	result := strings.Split(string(rawOutput), "\n")
	return result, nil
}
