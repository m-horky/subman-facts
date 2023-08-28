package facts

import (
	"bytes"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"os"
	"os/exec"
	"strings"
)

func CollectAll() map[string]any {
	everything := make(map[string]any)

	customCollector := NewCustomCollector()
	customFacts, err := customCollector.GetData(true)
	if err == nil {
		everything["custom"] = customFacts
	}

	distributionCollector := NewDistributionCollector()
	distributionFacts, err := distributionCollector.GetData(true)
	if err == nil {
		everything["distribution"] = distributionFacts
	}

	insightsCollector := NewInsightsCollector()
	insightsFacts, err := insightsCollector.GetData(true)
	if err == nil {
		everything["insights"] = insightsFacts
	}

	memoryCollector := NewMemoryCollector()
	memoryFacts, err := memoryCollector.GetData(true)
	if err == nil {
		everything["memory"] = memoryFacts
	}

	networkCollector := NewNetworkCollector()
	networkFacts, err := networkCollector.GetData(true)
	if err == nil {
		everything["network"] = networkFacts
	}

	systemCollector := NewSystemCollector()
	systemFacts, err := systemCollector.GetData(true)
	if err == nil {
		everything["system"] = systemFacts
	}

	unameCollector := NewUnameCollector()
	unameFacts, err := unameCollector.GetData(true)
	if err == nil {
		everything["uname"] = unameFacts
	}

	uptimeCollector := NewUptimeCollector()
	uptimeFacts, err := uptimeCollector.GetData(true)
	if err == nil {
		everything["uptime"] = uptimeFacts
	}

	// Following facts are only collectable with root permissions
	dmidecodeCollector := NewDmidecodeCollector()
	dmidecodeFacts, err := dmidecodeCollector.GetData(true)
	if err == nil {
		everything["dmidecode"] = dmidecodeFacts
	}

	virtCollector := NewVirtCollector()
	virtFacts, err := virtCollector.GetData(true)
	if err == nil {
		everything["virt"] = virtFacts
	}

	// Following facts are only collected when their binaries are installed
	kpatchCollector := NewKpatchCollector()
	kpatchFacts, err := kpatchCollector.GetData(true)
	if err == nil {
		everything["kpatch"] = kpatchFacts
	}

	// Following facts are only collectable in their specific cloud environments
	awsCollector := NewAWSCollector()
	awsFacts, err := awsCollector.GetData(true)
	if err == nil {
		everything["aws"] = awsFacts
	}
	azureCollector := NewAzureCollector()
	azureFacts, err := azureCollector.GetData(true)
	if err == nil {
		everything["azure"] = azureFacts
	}
	gcpCollector := NewGCPCollector()
	gcpFacts, err := gcpCollector.GetData(true)
	if err == nil {
		everything["gcp"] = gcpFacts
	}

	return everything
}

// CollectAllAsSubscriptionManager takes the output of CollectAll and converts it into
// data that are compatible with subscription-manager output.
// This includes omitting some values (MAC address of loopback interface) or transforming them
// (system UUID).
func CollectAllAsSubscriptionManager() map[string]any {
	// everything := CollectAll()

	// MAC of `lo` is omitted
	// DMI's UUID is uppercase
	// DMI's bios address is lowercase

	return make(map[string]any)
}

type OperatingSystem struct {
	Run  func(command string, args ...string) ([]string, []string, error)
	Read func(string) ([]string, error)
}

var OS = OperatingSystem{Run: run, Read: read}

// run invokes a shell program and returns the output as lines.
func run(command string, args ...string) ([]string, []string, error) {
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

// read reads a file and returns the lines.
func read(path string) ([]string, error) {
	rawOutput, err := os.ReadFile(path)
	if err != nil {
		return make([]string, 0), err
	}
	result := strings.Split(string(rawOutput), "\n")
	return result, nil
}

// TODO Probably move to internal

// Flatten takes in a nested map and returns it as a '.'-joined key-value map.
func Flatten(mappings map[string]any) map[string]string {
	result := make(map[string]string)

	for key, value := range mappings {
		switch value.(type) {
		case string:
			result[key] = value.(string)
		case int:
			result[key] = fmt.Sprintf("%d", value.(int))
		case float64:
			result[key] = fmt.Sprintf("%f", value.(float64))
		case map[string]any:
			for k, v := range Flatten(value.(map[string]any)) {
				result[fmt.Sprintf("%s.%s", key, k)] = fmt.Sprintf("%v", v)
			}
		default:
			log.Debugf("cannot flatten %#v=%#v", key, value)
		}
	}
	return result
}
