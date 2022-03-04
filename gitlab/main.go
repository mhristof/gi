package gitlab

import "errors"

type Client struct{}

func (c Client) Valid(remote string) bool {
	return false
}

var ErrNotImplemented = errors.New("not implemented")

func (c Client) WebURL(path string) (string, error) {
	return "", ErrNotImplemented
}
