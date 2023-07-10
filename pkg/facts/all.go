package facts

import (
	"fmt"
	"io"
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
func getCommandOutput(command string, args ...string) ([]string, []string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return []string{}, []string{}, err
	}
	cmdStr := path
	if len(args) > 0 {
		cmdStr = fmt.Sprintf("%s %s", path, strings.Join(args, " "))
	}

	cmd := exec.Command(path, args...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("could not connect stdout of %s: %s", path, err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("could not connect stderr of %s: %s", path, err)
	}

	err = cmd.Start()
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("could not invoke %s: %s", cmdStr, err)
	}

	stdout, _ := io.ReadAll(stdoutPipe)
	stderr, _ := io.ReadAll(stderrPipe)

	err = cmd.Wait()
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("could not run %s: %s", cmdStr, err)
	}

	return strings.Split(string(stdout), "\n"), strings.Split(string(stderr), "\n"), nil
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
