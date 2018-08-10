package pullrequests

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/slomek/gh/git"
	"github.com/slomek/gh/github"
	"github.com/spf13/cobra"
)

func init() {
	create.PersistentFlags().StringArray("assign", nil, "assignees for review")

	Command.AddCommand(create)
}

// Command groups all pr-related actions.
var Command = &cobra.Command{
	Use:   "pr",
	Short: "Manage PRs",
	Long:  "Manage GitHub's pull-requests",
}

// create creates a new GitHub pull request.
var create = &cobra.Command{
	Use:   "create",
	Short: "Create PR",
	Long:  "Create a new pull-request",
	Run: func(cmd *cobra.Command, args []string) {
		repoName, err := git.RepoName()
		if err != nil {
			fmt.Printf("Failed to get repo name: %v\n", err)
			return
		}

		message, err := git.CommitMessage()
		if err != nil {
			fmt.Printf("Failed to get commit message: %v\n", err)
			return
		}

		branch, err := git.BranchName()
		if err != nil {
			fmt.Printf("Failed to get branch name: %v\n", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		ghcli := github.NewClient(os.Getenv("GITHUB_PR_TOKEN"))
		pr, err := ghcli.CreatePullRequest(ctx, repoName, message, branch)
		if err != nil {
			fmt.Printf("Failed to create PR: %v", err)
			return
		}

		fmt.Printf("Created new pull request (#%d)\n", pr.Number)
		fmt.Printf("Details -> %s\n", *pr.HTMLURL)

		aa, err := cmd.Flags().GetStringArray("assign")
		if err != nil {
			fmt.Printf("Failer to read assignee list: %v", err)
			return
		}
		if len(aa) != 0 {
			ghcli.RequestReview(ctx, repoName, *pr.Number, aa)
		}
	},
}
