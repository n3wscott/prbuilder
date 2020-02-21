package commands

import (
	"log"

	"github.com/n3wscott/prbuilder/pkg/builder"
	"github.com/n3wscott/prbuilder/pkg/commands/options"
	"github.com/spf13/cobra"
)

var (
	gito = &options.GitOptions{}
	pro  = &options.PullRequestOptions{}
	//do   = &options.DryRunOptions{}
	vo = &options.VerboseOptions{}
)

func TopLevelRunE(cmd *cobra.Command, args []string) error {
	// Build up command.
	i := &builder.Builder{
		//DryRun:  do.DryRun,
		Verbose: vo.Verbose,

		// Git options.
		Workspace: gito.Workspace,
		Owner:     gito.Owner,
		Repo:      gito.Repo,
		Branch:    gito.Branch,

		// PR options.
		Title:     pro.Title,
		Body:      pro.Body,
		Token:     pro.Token,
		Signature: pro.Signature(),
	}

	// Run it.
	if err := i.Do(); err != nil {
		log.Fatalf("failed to run pr builder command: %v", err)
	}
	return nil
}

func AddCommands(topLevel *cobra.Command) {
	//options.AddDryRunArg(topLevel, do)
	options.AddGitArgs(topLevel, gito)
	options.AddPullRequestArgs(topLevel, pro)
	options.AddVerboseArg(topLevel, vo)

	addHelp(topLevel)
}
