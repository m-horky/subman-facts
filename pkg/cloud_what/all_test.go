package cloud_what

import (
	"fmt"
	"os"
)

func readTestFile(name string) string {
	raw, err := os.ReadFile(fmt.Sprintf("test_data/%s", name))
	if err != nil {
		panic(fmt.Sprintf("could not open test file %s: %s", name, err))
	}
	return string(raw)
}
