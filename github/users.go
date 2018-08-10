package github

import (
	"context"

	ghcli "github.com/google/go-github/github"
	"github.com/pkg/errors"
)

type OrganizationMember struct {
	Login string `json:"login,omitempty"`
}

func (c *Client) ListOrganizationUsers(ctx context.Context, orgName string) ([]*ghcli.User, error) {
	uu, _, err := c.client.Organizations.ListMembers(ctx, orgName, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list users in organization '%s'", orgName)
	}
	return uu, nil
}
