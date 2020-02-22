package commands

import (
	"fmt"
	"log"

	"github.com/n3wscott/prbuilder/pkg/builder"
	"github.com/n3wscott/prbuilder/pkg/commands/options"
	"github.com/spf13/cobra"
)

var (
	pro = &options.PullRequestOptions{}
	fso = &options.FileSystemOptions{}
	//do   = &options.DryRunOptions{}
	vo = &options.VerboseOptions{}
)

func NewTopLevelCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "prbuilder",
		Short:   "Make building GitHub Pull Requests automated.",
		RunE:    TopLevelRunE,
		Example: example(),
	}
	AddCommands(cmd)
	return cmd
}

func example() string {
	example1 := fmt.Sprintf(`
prbuilder \
  --workspace=./ \
  --repo=n3wscott/prbuilder \
  --target=master \
  --title="Doing a demo." \
  --body="Demo, ignore."
`)
	example2 := fmt.Sprintf(`
prbuilder \
  --workspace=./ \
  --repo=n3wscott/prbuilder \
  --target=master \
  --title="Fix spelling errors" \
  --body="Produced via: github.com/client9/misspell" \
  --name="Demo Person" \
  --email=demo@example.com
`)
	_ = example2

	return example1
}

// TopLevelRunE
func TopLevelRunE(cmd *cobra.Command, args []string) error {
	signature, err := pro.Signature()
	if err != nil {
		return err
	}

	token, _ := pro.Token()

	// Build up command.
	i := &builder.Builder{
		//DryRun:  do.DryRun,
		Verbose: vo.Verbose,

		// Auth
		Username: pro.Username(),
		Password: pro.Password(),

		// Filesystem
		Workspace: fso.Workspace,

		// Git options.
		Owner:        pro.Owner(),
		Repo:         pro.Repo(),
		Branch:       pro.BaseBranch,
		CommitBranch: pro.CommitBranch(),

		// PR options.
		Title: pro.Title,
		Body:  pro.Body,
		Token: token,

		// Author
		Signature: signature,
		Signoff:   pro.Signoff,
	}

	// Run it.
	if err := i.Do(); err != nil {
		log.Fatalf("failed to run pr builder command: %v", err)
	}
	return nil
}

func AddCommands(topLevel *cobra.Command) {

	//options.AddDryRunArg(topLevel, do)
	options.AddFileSystemArgs(topLevel, fso)
	options.AddPullRequestArgs(topLevel, pro)
	options.AddVerboseArg(topLevel, vo)

	addHelp(topLevel)
}
