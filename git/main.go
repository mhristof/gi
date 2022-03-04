package git

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/mhristof/gi/github"
	"github.com/mhristof/gi/gitlab"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Repo holds information about a repository.
type Repo struct {
	Git    *git.Repository
	Dir    string
	Client API
}

type API interface {
	// WebURL return the web URL of the given object
	WebURL(path, branch string, line int) (string, error)
	// Valid Return true if the input remote is a valid remote for this client (ie github, gitlab, etc)
	Valid(string) bool
}

// New Create a new git repository object from the given directory.
// The directory could be relative or absolute folder or file inside the git
// repository.
func New(directory string) (*Repo, error) {
	absDir, err := findGitFolder(directory)
	if err != nil {
		return nil, errors.Wrap(err, "Canot find .git folder in "+directory)
	}

	repo, err := git.PlainOpen(absDir)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open git repo")
	}

	config, err := repo.Config()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get config")
	}

	var client API

	remote := config.Remotes["origin"].URLs[0]
	clients := []API{
		gitlab.Client{Remote: remote},
		github.Client{Remote: remote},
	}

	for _, cl := range clients {
		if cl.Valid(remote) {
			client = cl

			break
		}
	}

	if client == nil {
		return nil, errors.New("cannot handle remote")
	}

	ret := &Repo{
		Git:    repo,
		Client: client,
		Dir:    absDir,
	}

	log.WithFields(log.Fields{
		"ret": ret,
	}).Debug("created git object")

	return ret, nil
}

// ErrNotAGitRepo is thrown when the given folder/config is not a git repository.
var ErrNotAGitRepo = errors.New("not a git repository")

func findGitFolder(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, "cannot find abs path")
	}

	parts := strings.Split(abs, "/")
	for i := len(parts); i > 0; i-- {
		thisPath := "/" + filepath.Join(parts[0:i]...)
		thisPathGit := filepath.Join(thisPath, ".git")

		if info, err := os.Stat(thisPathGit); err == nil && info.IsDir() {
			return thisPath, nil
		}
	}

	return "", ErrNotAGitRepo
}

func (r *Repo) WebURL(item string, line int) (string, error) {
	branch, err := r.Git.Head()
	if err != nil {
		return "", errors.Wrap(err, "cannot get branch")
	}
	branchName := branch.Name().String()

	ret, err := r.Client.WebURL(item, path.Base(branchName), line)
	if err != nil {
		return "", errors.Wrap(err, "cannot get URL")
	}

	return ret, nil
}

// // Branch Return the current branch of the git repository by reading .git/HEAD
// func (r *Repo) Branch() string {
// 	head, err := ioutil.ReadFile(filepath.Join(r.Dir, ".git/HEAD"))
// 	if err != nil {
// 		log.WithFields(log.Fields{
// 			"err":   err,
// 			"r.Dir": r.Dir,
// 		}).Panic("Cannot find .git/HEAD")
// 	}

// 	headS := strings.Split(strings.Split(string(head), "\n")[0], " ")[1]
// 	return headS
// }

// type Remote interface {
// 	File(string, string, int) (string, error)
// }

// // URL Returns the web url for the given file. Currently gitlab, github and codecommit
// // are supported
// func (r *Repo) URL(file string, line int) (string, error) {
// 	absFile, err := filepath.Abs(file)
// 	if err != nil {
// 		log.WithFields(log.Fields{
// 			"err":  err,
// 			"file": file,
// 		}).Panic("Cannot calculate abs path")
// 	}

// 	relativeFile := strings.TrimPrefix(strings.Replace(absFile, r.Dir, "", -1), "/")

// 	remotes := []Remote{
// 		&gitlab.Remote{R: r.Remote},
// 		&codecommit.Remote{R: r.Remote},
// 		&github.Remote{R: r.Remote},
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(remotes))
// 	res := make(chan string, 1)

// 	for _, remote := range remotes {
// 		go getURL(&wg, r, remote, relativeFile, line, res)
// 	}
// 	wg.Wait()

// 	return <-res, nil
// }

// func getURL(wg *sync.WaitGroup, r *Repo, remote Remote, relativeFile string, line int, c chan string) {
// 	defer wg.Done()

// 	this, err := remote.File(r.Branch(), relativeFile, line)
// 	if err != nil {
// 		return
// 	}
// 	c <- this
// }

// func redact(in string) string {
// 	token := regexp.MustCompile(`glpat-\w*`)

// 	return token.ReplaceAllString(in, "glpat-xxxxxx")
// }
