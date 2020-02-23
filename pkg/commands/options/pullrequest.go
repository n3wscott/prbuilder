package options

import (
	"strings"
	"time"

	"github.com/google/uuid"
	client "github.com/mattmoor/bindings/pkg/github"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// PullRequestOptions
type PullRequestOptions struct {

	// TODO: GitHub terms are base repo foo/bar @ base branch <- head repo:fork/bar compare branch
	// Git repos
	BaseRepo     string
	BaseBranch   string
	commitBranch string

	// Pull Request
	Title string
	Body  string
	token string

	// Signature
	name    string
	email   string
	Signoff bool
}

func AddPullRequestArgs(cmd *cobra.Command, o *PullRequestOptions) {

	// prbuilder --repo foo/bar --target master --branch=demo-branch -S --author "A U Thor <author@example.com>"

	// Git Repo

	cmd.Flags().StringVar(&o.BaseRepo, "repo", "",
		"The Github base <owner>/<repo> this PR will target. Required.")
	cmd.Flags().StringVar(&o.BaseBranch, "target", "",
		"The Github base <branch> name this PR will target. Required.")

	//cmd.Flags().StringVar(&o.CompareRepo, "compare-repo", "",
	//	"The Github modified <owner>/<repo> to create the PR from.")
	cmd.Flags().StringVarP(&o.commitBranch, "commit-branch", "b", "",
		"If provided, the branch to create based on the local changes. If value is 'random' a uuid will be used.")

	//cmd.Flags().StringVar(&o.Owner, "organization", "",
	//	"The Github organization to which we're sending a PR.")
	//cmd.Flags().StringVar(&o.Repo, "repository", "",
	//	"The Github repository to which we're sending a PR.")
	//cmd.Flags().StringVar(&o.Branch, "branch", "",
	//	"The branch we are building a PR against.")
	//cmd.Flags().StringVar(&o.PRBranch, "prbranch", "",
	//	"The branch that is created for this PR.")

	//_ = cmd.MarkFlagRequired("organization")
	//_ = cmd.MarkFlagRequired("repository")
	//_ = cmd.MarkFlagRequired("branch")
	//_ = cmd.MarkFlagRequired("prbranch")

	// Pull Request

	cmd.Flags().StringVar(&o.Title, "title", "",
		"The title of the PR to send.")
	cmd.Flags().StringVar(&o.Body, "body", "",
		"The body of the PR to send.")

	cmd.Flags().BoolVarP(&o.Signoff, "Signoff", "S", false,
		wrap80(`Add Signed-off-by line by the committer at the end of the commit log message. The meaning of a signoff depends on the project, but it typically certifies that committer has the rights to submit this work under the same license and agrees to a Developer Certificate of Origin (see http://developercertificate.org/ for more information).`))

	_ = cmd.MarkFlagRequired("title")
	_ = cmd.MarkFlagRequired("body")

	// Signature

	cmd.Flags().StringVar(&o.name, "name", "",
		"The author name.")
	cmd.Flags().StringVar(&o.email, "email", "",
		"The author email.")
	cmd.Flags().StringVar(&o.token, "token", "",
		"The random token for identifying this PR's provenance.")
}

func must(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

func (o *PullRequestOptions) Username() string {
	// TODO: we could pass these in the cli too?
	return must(client.ReadKey("username"))
}

func (o *PullRequestOptions) Password() string {
	// TODO: we could pass these in the cli too?
	return must(client.AccessToken())
}

// Owner returns the parsed owner from --repo.
func (o *PullRequestOptions) Owner() string {
	parts := strings.Split(o.BaseRepo, "/")
	if len(parts) == 2 {
		return parts[0]
	}
	panic("Expected repo to be <owner>/<repo>")
}

// Repo returns the parsed repo from --repo.
func (o *PullRequestOptions) Repo() string {
	parts := strings.Split(o.BaseRepo, "/")
	if len(parts) == 2 {
		return parts[1]
	}
	panic("Expected repo to be <owner>/<repo>")
}

// CommitBranch returns a pointer to a string to use as the commit branch name.
func (o *PullRequestOptions) CommitBranch() *string {
	if o.commitBranch == "" {
		return nil
	}
	if o.commitBranch == "random" {
		id, _ := uuid.NewUUID()
		o.commitBranch = id.String()
	}
	return &o.commitBranch
}

// Token returns the GitHub passed in command line, or from the local git config.
func (o *PullRequestOptions) Token() (*string, error) {
	if o.token == "" {
		token, err := gitconfig.GithubToken()
		if err != nil {
			return nil, err
		}
		o.token = token
	}
	return &o.token, nil
}

// Signature returns the structured git Signature if passed in from the username and email, or from local git config if set.
func (o *PullRequestOptions) Signature() (*object.Signature, error) {
	if o.name == "" {
		name, err := gitconfig.Username()
		if err != nil {
			return nil, err
		}
		o.name = name
	}

	if o.email == "" {
		email, err := gitconfig.Email()
		if err != nil {
			return nil, err
		}
		o.email = email
	}

	return &object.Signature{
		Name:  o.name,
		Email: o.email,
		When:  time.Now(),
	}, nil
}
