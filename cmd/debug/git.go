package debug

import (
	"fmt"

	"github.com/slomek/gh/git"
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(gitinfo)
}

// Command is used for debugging purposes.
var Command = &cobra.Command{
	Use:   "debug",
	Short: "Debug",
	Long:  "Debug",
}

// gitinfo prints information read from git metadata.
var gitinfo = &cobra.Command{
	Use:   "git",
	Short: "Git info",
	Long:  "Information read from git metadata",
	Run: func(cmd *cobra.Command, args []string) {
		repoName, err := git.RepoName()
		if err != nil {
			fmt.Printf("Repository name: failed to read: %v\n", err)
		} else {
			fmt.Printf("Repository name: %s\n", repoName)
		}

		message, err := git.CommitMessage()
		if err != nil {
			fmt.Printf("Commit message: failed to read: %v\n", err)
		} else {
			fmt.Printf("Commit message: %s\n", message)
		}

		branch, err := git.BranchName()
		if err != nil {
			fmt.Printf("Branch name: failed to read: %v\n", err)
		} else {
			fmt.Printf("Branch name: %s\n", branch)
		}
	},
}
