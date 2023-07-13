package facts

import (
	"fmt"
	"os/exec"
)

var directoryWithInstalledModules = "/var/lib/kpatch/"

// directoriesWithLoadedModules contains paths in which the modules could be found
// in various versions of kpatch.
var directoriesWithLoadedModules = []string{
	"/sys/kernel/livepatch/",
	"/sys/kernel/kpatch/patches/",
	"/sys/kernel/kpatch/",
}

type KpatchFacts struct {
	Installed string
	Loaded    string
}

type KpatchCollector struct {
	data      KpatchFacts
	collected bool
}

func NewKpatchCollector() KpatchCollector {
	return KpatchCollector{}
}

func (c *KpatchCollector) GetData(rescan bool) (KpatchFacts, error) {
	if rescan || !c.collected {
		c.data = KpatchFacts{}
	}

	err := c.collect()
	if err != nil {
		return KpatchFacts{}, err
	}
	return c.data, nil
}

func (c *KpatchCollector) collect() error {
	if !c.isInstalled() {
		return fmt.Errorf("kpatch binary is not installed")
	}

	return fmt.Errorf("Kpatch collection is not implemented.")
}

func (c *KpatchCollector) isInstalled() bool {
	_, err := exec.LookPath("kpatch")
	return err == nil
}
