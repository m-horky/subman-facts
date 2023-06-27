package collector

import (
	"os"
	"os/exec"
	"strings"
)

// Collector is an object which is able to collect some kind of facts about its
// host system.
type Collector interface {
	// GetData ensures the Collector has collected its facts.
	GetData() (map[string]string, error)
	// Flush ensures Collector has deleted previously collected data, if any.
	Flush()
}

// FIXME Currently, the collectors have no way of specifying defaults for some
//  key-value pair. When they do, they must do so at start of '.collect()'.

// getCommandOutput invokes a shell program and returns the output.
// It strips out the (last) newline.
func getCommandOutput(cmd string, args ...string) (string, error) {
	rawOutput, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(rawOutput), "\n"), nil
}

// getFileOutput reads a file and returns a slice of lines in it
func getFileOutput(path string) ([]string, error) {
	rawOutput, err := os.ReadFile(path)
	if err != nil {
		return make([]string, 0), err
	}
	result := strings.Split(string(rawOutput), "\n")
	return result, nil
}
