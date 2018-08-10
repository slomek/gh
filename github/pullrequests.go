package github

import (
	"context"
	"strings"

	ghcli "github.com/google/go-github/github"
	"github.com/pkg/errors"
)

func (c *Client) CreatePullRequest(ctx context.Context, repoWithOwner, title, head string) (*ghcli.PullRequest, error) {
	rwop := strings.Split(repoWithOwner, "/")
	owner := rwop[0]
	repo := rwop[1]

	base := "master"

	pr, _, err := c.client.PullRequests.Create(ctx, owner, repo, &ghcli.NewPullRequest{
		Title: &title,
		Head:  &head,
		Base:  &base,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create pull request")
	}
	return pr, nil
}
