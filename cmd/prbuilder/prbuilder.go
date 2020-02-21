package main

import (
	"log"

	"github.com/n3wscott/prbuilder/pkg/commands"
)

func main() {
	cmd := commands.NewTopLevelCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}