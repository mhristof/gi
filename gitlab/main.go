package gitlab

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Client struct {
	Remote string
}

func (c Client) Valid(remote string) bool {
	return strings.Contains(remote, "gitlab.com")
}

var ErrNotImplemented = errors.New("not implemented")

func (c Client) WebURL(path, branch string, line int) (string, error) {
	url := ssh2https(c.Remote)

	ret := fmt.Sprintf("%s/-/blob/%s/%s", url, branch, path)

	if line >= 0 {
		ret = fmt.Sprintf("%s#L%d", ret, line)
	}

	return ret, nil
}

func ssh2https(remote string) string {
	remote = strings.TrimPrefix(remote, "git@")
	remote = strings.TrimSuffix(remote, ".git")
	remote = strings.ReplaceAll(remote, "com:", "com/")

	if !strings.HasPrefix(remote, "http") {
		remote = "https://" + remote
	}

	if strings.Count(remote, ":") > 1 {
		// probably a user/pass is embedded in the url
		u, err := url.Parse(remote)
		if err != nil {
			return remote
		}

		remote = fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)
	}

	return remote
}
