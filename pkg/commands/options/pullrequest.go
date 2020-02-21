package options

import (
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// PullRequestOptions
type PullRequestOptions struct {
	Title string
	Body  string
	Token string

	name  string
	email string
}

func AddPullRequestArgs(cmd *cobra.Command, o *PullRequestOptions) {
	cmd.Flags().StringVar(&o.Title, "title", "",
		"The title of the PR to send.")
	cmd.Flags().StringVar(&o.Body, "body", "",
		"The body of the PR to send.")
	cmd.Flags().StringVar(&o.Token, "token", "",
		"The random token for identifying this PR's provenance.")

	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("body")
	_ = cmd.MarkFlagRequired("token")

	cmd.Flags().StringVar(&o.name, "name", "",
		"The author name.")
	cmd.Flags().StringVar(&o.email, "email", "",
		"The author email.")

	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("email")
}

func (o *PullRequestOptions) Signature() object.Signature {
	return object.Signature{
		Name:  o.name,
		Email: o.email,
		When:  time.Now(),
	}
}
