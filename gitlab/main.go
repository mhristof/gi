package gitlab

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mhristof/gi/keychain"
	"github.com/pkg/errors"
	glab "github.com/xanzy/go-gitlab"
)

type Client struct {
	Remote string
	api    *glab.Client
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

func New(remote string) Client {
	api, err := glab.NewClient(
		keychain.Item("GITLAB_READONLY_TOKEN"),
		glab.WithBaseURL("https://gitlab.com/api/v4"),
	)
	if err != nil {
		panic(err)
	}

	return Client{
		Remote: ssh2https(remote),
		api:    api,
	}
}

func (c Client) TransformReviewers(emails map[string]int) map[string]int {
	ret := map[string]int{}

	for email, v := range emails {
		ret[strings.Split(email, "@")[0]] = v
	}

	return ret
}
