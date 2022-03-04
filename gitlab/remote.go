package gitlab

// // ErrNotGitlab The remote doesnt seem like a Gitlab server.
// var ErrNotGitlab = errors.New("not a valid Gitlab remote")

// // Remote Represent a gitlab remote.
// type Remote struct {
// 	R string
// }

// // Valid Checks a remote to see if its a valid gitlab instance.
// func (r *Remote) Valid() bool {
// 	gitlabURL := regexp.MustCompile(`gitlab`)

// 	return gitlabURL.MatchString(r.R)
// }

// // URL Return the URL of the remote by sanitizing it.
// func (r *Remote) URL() string {
// 	remote := gitlabHTTP(r.R)
// 	if remote != "" {
// 		return remote
// 	}

// 	remote = gitlabSSH(r.R)
// 	if remote != "" {
// 		return remote
// 	}

// 	log.WithFields(log.Fields{
// 		"r.R": r.R,
// 	}).Error("Not a gitlab remote")

// 	return ""
// }

// func gitlabSSH(url string) string {
// 	if !strings.HasPrefix(url, "git@gitlab.com:") {
// 		return ""
// 	}

// 	url = strings.TrimSuffix(url, ".git")

// 	return strings.Replace(url, "git@gitlab.com:", "https://gitlab.com/", 1)
// }

// func gitlabHTTP(url string) string {
// 	remRegex := regexp.MustCompile(`https://(?P<username>.*):(?P<token>.*)@(?P<url>.*)`)
// 	match := remRegex.FindStringSubmatch(url)

// 	if remRegex.MatchString(url) {
// 		for i, name := range remRegex.SubexpNames() {
// 			if name == "url" {
// 				return fmt.Sprintf("https://%s", match[i])
// 			}
// 		}
// 	}

// 	return ""
// }

// // File Retrieves the file url for the given file. Throws a ErrNotGitlab
// // if the repository is not a valid gitlab url.
// func (r *Remote) File(branch, file string, line int) (string, error) {
// 	if !r.Valid() {
// 		return "", ErrNotGitlab
// 	}

// 	branch = strings.ReplaceAll(branch, "refs/heads/", "")
// 	ret := fmt.Sprintf("%s/-/blob/%s/%s", r.URL(), branch, file)

// 	if line >= 0 {
// 		ret += fmt.Sprintf("#L%d", line)
// 	}

// 	return ret, nil
// }
