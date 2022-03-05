package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mhristof/gi/util"
	"github.com/stretchr/testify/assert"
)

func NewGit(t *testing.T, commands []string) (string, func()) {
	t.Helper()

	dir, err := ioutil.TempDir("", "git")
	if err != nil {
		t.Fatal(err)
	}

	script(t, append([]string{
		"cd " + dir,
		"git init",
		"git remote add origin https://github.com/user/repo.git",
		"touch hi",
		"git add hi",
		"git commit -m 'hi'",
	}, commands...))

	return dir, func() {
		os.RemoveAll(dir)
	}
}

func script(t *testing.T, commands []string) {
	t.Helper()

	_, _, err := util.Bash(strings.Join(commands, " && "))
	if err != nil {
		t.Fatal(err)
	}
}

func TestBranchFilesChanged(t *testing.T) {
	cases := []struct {
		name     string
		setup    []string
		expected []string
	}{
		{
			name: "simple case",
			setup: []string{
				"touch c && git add c && git commit -am 'add c'",
				"touch d && git add d && git commit -am 'add d'",
				"git checkout -b testing1",
				"touch a && git add a && git commit -am 'add a'",
				"touch b && git add b && git commit -am 'add b'",
				"date >> d && git add d && git commit -am 'update d'",
			},
			expected: []string{"a", "b", "d"},
		},
	}

	for _, test := range cases {
		dir, _ := NewGit(t, test.setup)

		gg, err := New(dir)
		assert.Nil(t, err, test.name)

		files, err := gg.BranchFilesChanged("testing1", "main")
		assert.Nil(t, err, test.name)
		assert.Equal(t, test.expected, files, test.name)
		fmt.Println(fmt.Sprintf("dir: %+v %T", dir, dir))
	}
}
