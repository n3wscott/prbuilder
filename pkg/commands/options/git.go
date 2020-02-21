package options

import (
	"github.com/spf13/cobra"
)

// GitOptions
type GitOptions struct {
	Workspace string
	Owner     string
	Repo      string
	Branch    string
}

func AddGitArgs(cmd *cobra.Command, o *GitOptions) {
	cmd.Flags().StringVar(&o.Branch, "workspace", "",
		"The workspace directory to turn into a PR.")
	cmd.Flags().StringVar(&o.Owner, "organization", "",
		"The Github organization to which we're sending a PR.")
	cmd.Flags().StringVar(&o.Repo, "repository", "",
		"The Github repository to which we're sending a PR.")
	cmd.Flags().StringVar(&o.Branch, "branch", "",
		"The branch we are building a PR against.")

	_ = cmd.MarkFlagRequired("workspace")
	_ = cmd.MarkFlagRequired("organization")
	_ = cmd.MarkFlagRequired("repository")
	_ = cmd.MarkFlagRequired("branch")
}
