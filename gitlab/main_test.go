package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSsh2https(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "http url",
			in:   "https://gitlab.com/user/project",
			out:  "https://gitlab.com/user/project",
		},
		{
			name: "ssh url",
			in:   "git@gitlab.com:user/project.git",
			out:  "https://gitlab.com/user/project",
		},
		{
			name: "http url with username/token",
			in:   "https://user:token@gitlab.com/user/project",
			out:  "https://gitlab.com/user/project",
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, ssh2https(test.in), test.name)
	}
}
