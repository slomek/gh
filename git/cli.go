package git

import (
	"context"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// RepoName returns a Github repository name (owner/repo) from local directory.
func RepoName() (string, error) {
	cmd := exec.CommandContext(context.Background(), "git", "config", "--get", "remote.origin.url")
	bs, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "failed to get `git config` command output")
	}
	name := string(bs)
	name = strings.TrimSpace(name)
	name = strings.TrimPrefix(name, "git@github.com:")
	name = strings.TrimSuffix(name, ".git")
	return name, nil
}

// CommitMessage returns the last commit message from diff between current branch and master.
// If there are more than one commits, an empty string is being returned.
func CommitMessage() (string, error) {
	cmd := exec.CommandContext(context.Background(), "git", "log", "origin/master..HEAD", `--pretty=format:%s`)
	bs, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "failed to get `git log` command output")
	}
	raw := string(bs)
	message := strings.TrimSuffix(raw, "/n")

	commits := strings.Split(message, "\n")
	if len(commits) == 1 {
		return commits[0], nil
	}

	return "", errors.New("multiple or zero commits, sorry!")
}

// BranchName returns the name of a currently checked-out branch in the local directory.
func BranchName() (string, error) {
	cmd := exec.CommandContext(context.Background(), "git", "rev-parse", "--abbrev-ref", "HEAD")
	bs, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "failed to get `git log` command output")
	}
	raw := string(bs)
	return strings.TrimSpace(raw), nil
}
