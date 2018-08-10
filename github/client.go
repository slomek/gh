package github

import (
	"context"

	ghcli "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Client struct {
	client *ghcli.Client
}

func NewClient(token string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, ts)

	return &Client{
		client: ghcli.NewClient(httpClient),
	}
}
