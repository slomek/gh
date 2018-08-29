package reviews

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/slomek/gh/git"
	"github.com/slomek/gh/github"
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(request)
	Command.AddCommand(list)
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

		ghcli, err := github.NewClient(os.Getenv("GITHUB_PR_TOKEN"))
		if err != nil {
			fmt.Printf("Failed to create GitHub client: %v\n", err)
			return
		}

		err = ghcli.RequestReview(ctx, repoName, prNo, ppl)
		if err != nil {
			fmt.Printf("Failed to create PR: %v", err)
			return
		}
	},
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List review requests",
	Long:  "Lists all reviews assigned to you",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ghcli, err := github.NewClient(os.Getenv("GITHUB_PR_TOKEN"))
		if err != nil {
			fmt.Printf("Failed to create GitHub client: %v\n", err)
			return
		}

		issues, err := ghcli.ListIssuesWithRequestedReview(context.Background(), "repoName", 0, []string{"ppl"})
		if err != nil {
			fmt.Printf("Failed to list my code review requests: %v", err)
			return
		}

		fmt.Printf("There are %d PRs waiting for the review:\n", len(issues))
		for _, i := range issues {
			fmt.Printf(" - %v by %v %v\n", *i.HTMLURL, color.BlueString(*i.User.Login), color.YellowString("(%v ago)", formatTimeAgo(*i.CreatedAt)))
		}
	},
}

func formatTimeAgo(t time.Time) string {
	d := time.Since(t)
	if d.Hours() < 24 {
		return fmt.Sprintf("%.0fh", d.Hours())
	}
	if d.Hours() < 48 {
		return "a day"
	}
	if d.Hours() < 24*30 {
		return fmt.Sprintf("%.0f days", d.Hours()/24)
	}
	return fmt.Sprintf("%.0f months", d.Hours()/24/30)
}
