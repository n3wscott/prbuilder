package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/n3wscott/prbuilder/pkg/commands"
)

func main() {
	cmds := &cobra.Command{
		Use:   "prbuilder",
		Short: "TODO.",
		RunE:  commands.TopLevelRunE,
	}
	commands.AddCommands(cmds)

	if err := cmds.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
