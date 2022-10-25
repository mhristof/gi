package git

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/mhristof/gi/github"
	"github.com/mhristof/gi/gitlab"
	"github.com/mhristof/gi/util"
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
	TransformReviewers(map[string]int) map[string]int
}

// New Create a new git repository object from the given directory.
// The directory could be relative or absolute folder or file inside the git
// repository.
func New(directory string) (*Repo, error) {
	absDir, err := findGitFolder(directory)
	if err != nil {
		return nil, errors.Wrap(err, "Canot find .git folder in "+directory)
	}

	start := time.Now()
	repoD, err := git.PlainOpen(absDir)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot open [%s]", directory)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.WithFields(log.Fields{
		"elapsed": elapsed,
	}).Debug("git.PlainOpen")

	configD, err := repoD.Config()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get git config")
	}

	fs := memfs.New()
	storer := memory.NewStorage()

	log.WithFields(log.Fields{
		"absDir": absDir,
	}).Debug("cloning git")

	start = time.Now()
	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: absDir,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot open git repo")
	}

	t = time.Now()
	elapsed = t.Sub(start)
	log.WithFields(log.Fields{
		"elapsed": elapsed,
	}).Debug("git.Clone")

	opts := &git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
	}

	if err := repo.Fetch(opts); err != nil {
		return nil, errors.Wrap(err, "cannot fetch branches to inmem clone")
	}

	err = repo.SetConfig(configD)
	if err != nil {
		return nil, errors.Wrap(err, "cannot set inmem config")
	}

	config, err := repo.Config()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get config")
	}

	var client API

	remote := config.Remotes["origin"].URLs[0]
	clients := []API{
		gitlab.New(remote),
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
	branch, err := r.BranchName()
	if err != nil {
		return "", errors.Wrap(err, "cannot get branch")
	}

	ret, err := r.Client.WebURL(item, branch, line)
	if err != nil {
		return "", errors.Wrap(err, "cannot get URL")
	}

	return ret, nil
}

// Reviewers Returns a map of authors and the percentage of line changes they
// have in the current file.
func (r *Repo) Reviewers(item string) (map[string]int, error) {
	head, err := r.Git.Head()
	if err != nil {
		return map[string]int{}, errors.Wrap(err, "cannot get Head")
	}

	commit, err := r.Git.CommitObject(head.Hash())
	if err != nil {
		return map[string]int{}, errors.Wrap(err, "cannot get commit object")
	}

	blame, err := git.Blame(commit, item)
	if err != nil {
		mainName, err := r.Main()
		if err != nil {
			return map[string]int{}, errors.Wrap(err, "cannot get main branch name")
		}

		blame, err = r.BlameFromBranch(item, mainName)
		if err != nil {
			return map[string]int{}, errors.Wrap(err, "cannot blame")
		}
	}

	authors := map[string]int{}
	for _, line := range blame.Lines {
		authors[line.Author]++
	}

	for author, lines := range authors {
		authors[author] = (lines * 100) / len(blame.Lines)
	}

	log.WithFields(log.Fields{
		"authors": authors,
		"file":    item,
	}).Debug("git blame")
	return r.Client.TransformReviewers(authors), nil
}

func (r *Repo) BlameFromBranch(file, branch string) (*git.BlameResult, error) {
	branchR := r.Branch(branch)
	if branchR == nil {
		return nil, fmt.Errorf("cannot find branch [%s]", branch)
	}

	commit, err := r.Git.CommitObject(branchR.Hash())
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get commit object for [%s]", branch)
	}

	blame, err := git.Blame(commit, file)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot blame file [%s] from [%s]", file, branch)
	}

	return blame, nil
}

func (r *Repo) BranchName() (string, error) {
	branch, err := r.Git.Head()
	if err != nil {
		return "", errors.Wrap(err, "cannot get branch")
	}

	branchName := branch.Name().String()

	return path.Base(branchName), nil
}

func (r *Repo) Branch(name string) *plumbing.Reference {
	branches, err := r.Git.Branches()
	if err != nil {
		panic(err)
	}

	var branch *plumbing.Reference

	branches.ForEach(func(r *plumbing.Reference) error {
		if strings.Contains(r.String(), name) {
			branch = r
		}

		return nil
	})

	return branch
}

func (r *Repo) Main() (string, error) {
	for _, main := range []string{"main", "master"} {
		m := r.Branch(main)
		if m != nil {
			return main, nil
		}
	}

	return "", errors.New("cannot find main branch")
}

func (r *Repo) BranchReviewers() (map[string]int, error) {
	branch, err := r.BranchName()
	if err != nil {
		return map[string]int{}, errors.Wrap(err, "cannot get current branch")
	}

	main, err := r.Main()
	if err != nil {
		return map[string]int{}, errors.Wrap(err, " cannot find main branch")
	}

	files, err := r.BranchFilesChanged(branch, main)
	if err != nil {
		return map[string]int{}, errors.Wrap(err, "cannot calculate files changes")
	}

	authors := map[string]int{}
	for _, file := range files {
		fileAuthors, err := r.Reviewers(file)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Warning("cannot blame file")

			continue
		}

		log.WithFields(log.Fields{
			"authors": fileAuthors,
			"file":    file,
		}).Debug("git blame")

		authors = util.MapMerge(authors, fileAuthors)
	}

	return r.Client.TransformReviewers(authors), nil
}

func (r *Repo) BranchFilesChanged(src, dest string) ([]string, error) {
	tree, err := r.BranchTree(src)
	if err != nil {
		return []string{}, errors.Wrap(err, "cannot get src tree")
	}

	treeDest, err := r.BranchTree(dest)
	if err != nil {
		return []string{}, errors.Wrap(err, "cannot get dest tree")
	}

	changes, err := object.DiffTree(tree, treeDest)
	if err != nil {
		return []string{}, errors.Wrap(err, "cannot get diff from src and dest")
	}

	files := make([]string, len(changes))
	for i, change := range changes {
		files[i] = strings.TrimSuffix(strings.Fields(change.String())[3], ">")
	}

	return files, nil
}

func (r *Repo) BranchTree(name string) (*object.Tree, error) {
	branches, err := r.Git.Branches()
	if err != nil {
		panic(err)
	}

	var branch plumbing.ReferenceName

	_ = branches.ForEach(func(r *plumbing.Reference) error {
		if path.Base(r.Name().String()) == name {
			branch = r.Name()

			return nil
		}

		return nil
	})

	if branch.String() == "" {
		return nil, fmt.Errorf("branch not found [%s]", name)
	}

	ref, err := r.Git.Reference(branch, true)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get dest branch ref")
	}

	commit, err := r.Git.CommitObject(ref.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get dest branch commit")
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get dest branch tree")
	}

	return tree, nil
}
