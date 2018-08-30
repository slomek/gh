# gh - GitHub CLI

## Installation

```
go get -u github.com/slomek/gh
```

## Usage

### Pull requests

Creating a pull request from currently checked-out branch:
```
gh pr create
```

Creating a pull request with users "user1" and "user2" requested for a code review: 
```
gh pr create --assign user1 --assign user2
```

Requesting a review on an existing pull request (PR number 123, request review from user3 and user4):
```
gh reviews request 123 user3 user4
```

### Code reviews

Listing all pull requests, where you have been selected as a reviewer:
```
gh reviews list
```

### Listing users

In order to request a review, you need a login of the particular user from the organization. To list users from "my-organization", use:
```
gh users org my-organization
```

### Debugging

To see data read from git meta:
```
gh debug git
```
