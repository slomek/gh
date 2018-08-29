package github

import (
	"context"
	"sort"
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
func (c *Client) ListIssuesWithRequestedReview(ctx context.Context, repoWithOwner string, number int, revievers []string) ([]ghcli.Issue, error) {
	ii, _, err := c.client.Search.Issues(ctx, "is:open+is:pr+review-requested:slomek", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request review on pull request")
	}

	sort.Slice(ii.Issues, func(i, j int) bool {
		issI, issJ := ii.Issues[i], ii.Issues[j]
		return issI.CreatedAt.Before(*(issJ).CreatedAt)
	})

	return ii.Issues, nil
}
