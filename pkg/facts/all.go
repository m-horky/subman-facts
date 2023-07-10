package facts

import (
	"os"
	"os/exec"
	"strings"
)

type AllFacts struct {
	AWSFacts          AWSFacts          `json:"aws"`
	AzureFacts        AzureFacts        `json:"azure"`
	GCPFacts          GCPFacts          `json:"gcp"`
	CustomFacts       CustomFacts       `json:"custom"`
	DistributionFacts DistributionFacts `json:"distribution"`
	DmidecodeFacts    DmidecodeFacts    `json:"dmidecode"`
	InsightsFacts     InsightsFacts     `json:"insights"`
	KpatchFacts       KpatchFacts       `json:"kpatch"`
	MemoryFacts       MemoryFacts       `json:"memory"`
	NetworkFacts      NetworkFacts      `json:"network"`
	SystemFacts       SystemFacts       `json:"system"`
	UnameFacts        UnameFacts        `json:"uname"`
	UptimeFacts       UptimeFacts       `json:"uptime"`
	VirtFacts         VirtFacts         `json:"virt"`
}

func CollectAll() AllFacts {
	everything := AllFacts{}

	// TODO Capture the errors _somehow_?
	//  They are currently ignored; collection issues should already be logged,
	//  so the fact that the collection failed is already known to the user.
	everything.AWSFacts, _ = (&AWSCollector{}).GetData(true)
	everything.AzureFacts, _ = (&AzureCollector{}).GetData(true)
	everything.GCPFacts, _ = (&GCPCollector{}).GetData(true)
	everything.CustomFacts, _ = (&CustomCollector{}).GetData(true)
	everything.DistributionFacts, _ = (&DistributionCollector{}).GetData(true)
	everything.DmidecodeFacts, _ = (&DmidecodeCollector{}).GetData(true)
	everything.InsightsFacts, _ = (&InsightsCollector{}).GetData(true)
	everything.KpatchFacts, _ = (&KpatchCollector{}).GetData(true)
	everything.MemoryFacts, _ = (&MemoryCollector{}).GetData(true)
	everything.SystemFacts, _ = (&SystemCollector{}).GetData(true)
	everything.UnameFacts, _ = (&UnameCollector{}).GetData(true)
	everything.UptimeFacts, _ = (&UptimeCollector{}).GetData(true)
	everything.VirtFacts, _ = (&VirtCollector{}).GetData(true)

	return everything
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
