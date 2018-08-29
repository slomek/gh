package users

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/slomek/gh/github"
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(listFromOrganization)
}

// Command groups all user-related actions.
var Command = &cobra.Command{
	Use:   "users",
	Short: "List users",
	Long:  "List GitHub users",
}

// listFromOrganization lists users belonging to a given GitHub organization.
var listFromOrganization = &cobra.Command{
	Use:   "org",
	Short: "List users for organization",
	Long:  "List users for organization",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ghcli, err := github.NewClient(os.Getenv("GITHUB_PR_TOKEN"))
		if err != nil {
			fmt.Printf("Failed to create GitHub client: %v\n", err)
			return
		}
		orgName := args[0]

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		uu, err := ghcli.ListOrganizationUsers(ctx, orgName)
		if err != nil {
			fmt.Printf("Failed to list users: %v\n", err)
			return
		}

		fmt.Printf("Found %d members in organization %s:\n", len(uu), orgName)
		for _, u := range uu {
			fmt.Printf("- %s\n", *u.Login)
		}
	},
}
