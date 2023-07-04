package main

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~spc/go-log"
	"github.com/m-horky/subman-facts/pkg/project"
)

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

	fmt.Printf("subman-facts, version %s\n", project.Version)

	collectedFacts, errors := project.CollectAll()

	s, _ := json.MarshalIndent(collectedFacts, "", "\t")
	fmt.Println(string(s))

	for _, err := range errors {
		log.Error(err)
	}
}
