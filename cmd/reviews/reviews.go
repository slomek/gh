package reviews

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/slomek/gh/git"
	"github.com/slomek/gh/github"
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(request)
}

// Command groups all code review-related actions.
var Command = &cobra.Command{
	Use:   "reviews",
	Short: "Manage CRs",
	Long:  "Manage GitHub's code reviews",
}

// create requests a new GitHub code review.
var request = &cobra.Command{
	Use:   "request",
	Short: "Request CR",
	Long:  "Request a code review",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		repoName, err := git.RepoName()
		if err != nil {
			fmt.Printf("Failed to get repo name: %v\n", err)
			return
		}

		prNo, _ := strconv.Atoi(args[0])
		ppl := args[1:]

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		ghcli := github.NewClient(os.Getenv("GITHUB_PR_TOKEN"))
		err = ghcli.RequestReview(ctx, repoName, prNo, ppl)
		if err != nil {
			fmt.Printf("Failed to create PR: %v", err)
			return
		}
	},
}
