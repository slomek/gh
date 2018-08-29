package github

import (
	"context"

	ghcli "github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type Client struct {
	client *ghcli.Client
	myself string
}

func (c Client) MyLogin() string {
	return c.myself
}

func NewClient(token string) (*Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, ts)

	cli := ghcli.NewClient(httpClient)

	u, _, err := cli.Users.Get(context.Background(), "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current user's login")
	}

	return &Client{
		client: cli,
		myself: *u.Login,
	}, nil
}
