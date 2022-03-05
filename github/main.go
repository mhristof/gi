package github

import (
	"errors"
	"fmt"
	"strings"
)

type Client struct {
	Remote string
}

func (c Client) Valid(remote string) bool {
	return strings.Contains(remote, "github.com")
}

var ErrNotImplemented = errors.New("not implemented")

func (c Client) WebURL(path, branch string, line int) (string, error) {
	url := ssh2https(c.Remote)

	ret := fmt.Sprintf("%s/blob/%s/%s", url, branch, path)

	if line > 0 {
		ret = fmt.Sprintf("%s#L%d", ret, line)
	}

	return ret, nil
}

func ssh2https(remote string) string {
	remote = strings.TrimPrefix(remote, "git@")
	remote = strings.TrimSuffix(remote, ".git")
	remote = strings.ReplaceAll(remote, ":", "/")

	return "https://" + remote
}

func (c Client) TransformReviewers(emails map[string]int) map[string]int {
	return emails
}
