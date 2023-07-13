package facts

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type AllFacts struct {
	CustomFacts       CustomFacts       `json:"custom"`
	DistributionFacts DistributionFacts `json:"distribution"`
	InsightsFacts     InsightsFacts     `json:"insights"`
	MemoryFacts       MemoryFacts       `json:"memory"`
	NetworkFacts      NetworkFacts      `json:"network"`
	SystemFacts       SystemFacts       `json:"system"`
	UnameFacts        UnameFacts        `json:"uname"`
	UptimeFacts       UptimeFacts       `json:"uptime"`
	VirtFacts         VirtFacts         `json:"virt"`
	DmidecodeFacts    *DmidecodeFacts   `json:"dmidecode,omitempty"`
	KpatchFacts       *KpatchFacts      `json:"kpatch"`
	AWSFacts          *AWSFacts         `json:"aws,omitempty"`
	AzureFacts        *AzureFacts       `json:"azure,omitempty"`
	GCPFacts          *GCPFacts         `json:"gcp,omitempty"`
}

func CollectAll() AllFacts {
	everything := AllFacts{}

	customCollector := NewCustomCollector()
	everything.CustomFacts, _ = customCollector.GetData(true)
	distributionCollector := NewDistributionCollector()
	everything.DistributionFacts, _ = distributionCollector.GetData(true)
	insightsCollector := NewInsightsCollector()
	everything.InsightsFacts, _ = insightsCollector.GetData(true)
	memoryCollector := NewMemoryCollector()
	everything.MemoryFacts, _ = memoryCollector.GetData(true)
	networkCollector := NewNetworkCollector()
	everything.NetworkFacts, _ = networkCollector.GetData(true)
	systemCollector := NewSystemCollector()
	everything.SystemFacts, _ = systemCollector.GetData(true)
	unameCollector := NewUnameCollector()
	everything.UnameFacts, _ = unameCollector.GetData(true)
	uptimeCollector := NewUptimeCollector()
	everything.UptimeFacts, _ = uptimeCollector.GetData(true)
	virtCollector := NewVirtCollector()
	everything.VirtFacts, _ = virtCollector.GetData(true)

	// Following facts are only collectable with root permissions
	dmidecodeCollector := NewDmidecodeCollector()
	dmiFacts, err := dmidecodeCollector.GetData(true)
	if err == nil {
		*everything.DmidecodeFacts = dmiFacts
	}

	// Following facts are only collected when their binaries are installed
	kpatchCollector := NewKpatchCollector()
	kpatchFacts, err := kpatchCollector.GetData(true)
	if err == nil {
		*everything.KpatchFacts = kpatchFacts
	}

	// Following facts are only collectable in their specific cloud environments
	awsCollector := NewAWSCollector()
	awsFacts, err := awsCollector.GetData(true)
	if err == nil {
		*everything.AWSFacts = awsFacts
	}
	azureCollector := NewAzureCollector()
	azureFacts, err := azureCollector.GetData(true)
	if err == nil {
		*everything.AzureFacts = azureFacts
	}
	gcpCollector := NewGCPCollector()
	gcpFacts, err := gcpCollector.GetData(true)
	if err == nil {
		*everything.GCPFacts = gcpFacts
	}

	return everything
}

// getCommandOutput invokes a shell program and returns the output as lines.
func getCommandOutput(command string, args ...string) ([]string, []string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return []string{}, []string{}, err
	}
	cmdStr := path
	if len(args) > 0 {
		cmdStr = fmt.Sprintf("%s %s", path, strings.Join(args, " "))
	}

	// TODO Use context to set languages to C.UTF-8
	cmd := exec.Command(path, args...)
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err = cmd.Run()
	stdout := strings.Split(strings.TrimRight(outBuffer.String(), "\n"), "\n")
	stderr := strings.Split(strings.TrimRight(errBuffer.String(), "\n"), "\n")
	if err != nil {
		return stdout, stderr, fmt.Errorf("could not invoke %s: %s", cmdStr, err)
	}

	return stdout, stderr, nil
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
