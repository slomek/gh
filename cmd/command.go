package cmd

import (
	"github.com/slomek/gh/cmd/debug"
	"github.com/slomek/gh/cmd/pullrequests"
	"github.com/slomek/gh/cmd/reviews"
	"github.com/slomek/gh/cmd/users"

	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(pullrequests.Command)
	Command.AddCommand(reviews.Command)
	Command.AddCommand(users.Command)
	Command.AddCommand(debug.Command)
}

// Command is a root command of the application, it deoesn't do anything on its own.
var Command = &cobra.Command{
	Use:   "gh",
	Short: "CLI client for GitHub",
	Long:  "Command line too to make GitHub interaction faster",
}
