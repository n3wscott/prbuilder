# PR Builder

`prbuilder` is a tool for making prs to github.

[![GoDoc](https://godoc.org/github.com/n3wscott/prbuilder?status.svg)](https://godoc.org/github.com/n3wscott/prbuilder)
[![Go Report Card](https://goreportcard.com/badge/n3wscott/prbuilder)](https://goreportcard.com/report/n3wscott/prbuilder)

_Work in progress._

## Installation

`prbuilder` can be installed via:

```shell
go get github.com/n3wscott/prbuilder/cmd/prbuilder
```

To update your installation:

```shell
go get -u github.com/n3wscott/prbuilder/cmd/prbuilder
```

## Usage

`prbuilder` is only a base command. 

```shell
Usage:
  prbuilder [flags]

Flags:
      --body string           The body of the PR to send.
      --branch string         The branch we are building a PR against.
      --dry-run               Output what would happen.
      --email string          The author email.
      --name string           The author name.
      --organization string   The Github organization to which we're sending a PR.
      --repository string     The Github repository to which we're sending a PR.
      --title string          The title of the PR to send.
      --token string          The random token for identifying this PR's provenance.
  -V, --verbose               Output more debug info to stderr
      --workspace string      The workspace directory to turn into a PR.
```
