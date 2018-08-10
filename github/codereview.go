package github

import (
	"context"
	"strings"

	ghcli "github.com/google/go-github/github"
	"github.com/pkg/errors"
)

func (c *Client) RequestReview(ctx context.Context, repoWithOwner string, number int, revievers []string) error {
	rwop := strings.Split(repoWithOwner, "/")
	owner := rwop[0]
	repo := rwop[1]

	_, _, err := c.client.PullRequests.RequestReviewers(ctx, owner, repo, number, ghcli.ReviewersRequest{
		Reviewers: revievers,
	})
	if err != nil {
		return errors.Wrap(err, "failed to request review on pull request")
	}
	return nil
}
