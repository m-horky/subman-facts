package main

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/facts"
	"os"
)

const Version string = "upstream"

func configureLogging() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	level, err := log.ParseLevel("debug")
	if err != nil {
		level = log.LevelError
	}
	log.SetLevel(level)
}

func main() {
	configureLogging()

	_, _ = fmt.Fprintf(os.Stderr, "subman-facts, version %s\n", Version)

	collectedFacts, errors := facts.CollectAll()
	s, _ := json.MarshalIndent(collectedFacts, "", "\t")
	fmt.Println(string(s))

	for _, err := range errors {
		log.Error(err)
	}
}
