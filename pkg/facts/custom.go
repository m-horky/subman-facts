package facts

import (
	"encoding/json"
	"git.sr.ht/~spc/go-log"
	"os"
	"path"
	"strings"
)

type CustomCollector struct {
	data      map[string]any
	collected bool
	directory string
}

func NewCustomCollector() CustomCollector {
	return CustomCollector{
		data:      make(map[string]any),
		collected: false,
		directory: "/etc/rhsm/facts/",
	}
}

func (c *CustomCollector) GetData(rescan bool) (map[string]any, error) {
	if rescan || !c.collected {
		c.data = make(map[string]any)
	}

	err := c.collect()
	if err != nil {
		return make(map[string]any), err
	}
	return c.data, nil
}

func (c *CustomCollector) collect() error {
	// Custom facts can be stored as any JSON file in /etc/rhsm/facts/ directory.
	files, err := os.ReadDir(c.directory)
	if err != nil {
		log.Errorf("fact directory %s not found.", c.directory)
		return err
	}

	facts := make(map[string]any)

	for _, file := range files {
		filePath := path.Join(c.directory, file.Name())
		fileLines, err := OS.Read(filePath)
		if err != nil {
			log.Debugf("could not read fact file %s", filePath)
			continue
		}

		var fileFacts map[string]any
		err = json.Unmarshal([]byte(strings.Join(fileLines, " ")), &fileFacts)
		if err != nil {
			log.Warnf("could not parse fact file %s: %s", filePath, err)
			continue
		}

		for k, v := range fileFacts {
			if k == "custom" {
				log.Warnf("ignoring key '%s' in file %s: this fact cannot be overwritten", k, filePath)
				continue
			}
			if val, ok := facts[k]; ok {
				log.Warnf("overwriting key '%s = %s' with new value '%s' from file %s", k, val, v, filePath)
			}
			facts[k] = v
		}
	}

	c.data = facts
	return nil
}
