package codecommit

// func TestValid(t *testing.T) {
// 	var cases = []struct {
// 		name string
// 		r    string
// 		exp  bool
// 	}{
// 		{
// 			name: "simple codecommit repo",
// 			r:    "https://git-codecommit.region.amazonaws.com/v1/repos/foobar",
// 			exp:  true,
// 		},
// 		{
// 			name: "github repo",
// 			r:    "https://github.com/user/repo.git",
// 			exp:  false,
// 		},
// 	}

// 	for _, test := range cases {
// 		remote := Remote{R: test.r}
// 		assert.Equal(t, test.exp, remote.Valid(), test.name)
// 	}
// }

// func TestURL(t *testing.T) {
// 	var cases = []struct {
// 		name   string
// 		remote string
// 		exp    string
// 	}{
// 		{
// 			name:   "simple code commit repo",
// 			remote: "https://git-codecommit.region.amazonaws.com/v1/repos/foobar",
// 			exp:    "https://region.console.aws.amazon.com/codesuite/codecommit/repositories/foobar",
// 		},
// 	}

// 	for _, test := range cases {
// 		r := Remote{R: test.remote}

// 		assert.Equal(t, test.exp, r.URL(), test.name)

// 	}
// }

// func TestFile(t *testing.T) {
// 	var cases = []struct {
// 		name   string
// 		remote string
// 		branch string
// 		file   string
// 		line   int
// 		exp    string
// 	}{
// 		{
// 			name:   "simple file",
// 			remote: "https://git-codecommit.eu-west-2.amazonaws.com/v1/repos/repo",
// 			branch: "master",
// 			file:   "test/readme.md",
// 			line:   -1,
// 			exp:    "https://eu-west-2.console.aws.amazon.com/codesuite/codecommit/repositories/repo/browse/refs/heads/master/--/test/readme.md?region=eu-west-2",
// 		},
// 		{
// 			name:   "simple file",
// 			remote: "https://git-codecommit.eu-west-2.amazonaws.com/v1/repos/repo",
// 			branch: "master",
// 			file:   "test/readme.md",
// 			line:   100,
// 			exp:    "https://eu-west-2.console.aws.amazon.com/codesuite/codecommit/repositories/repo/browse/refs/heads/master/--/test/readme.md?region=eu-west-2#L100-100",
// 		},
// 	}

// 	for _, test := range cases {
// 		r := Remote{R: test.remote}

// 		url, err := r.File(test.branch, test.file, test.line)
// 		assert.Nil(t, err, test.name)
// 		assert.Equal(t, test.exp, url, test.name)
// 	}
// }
