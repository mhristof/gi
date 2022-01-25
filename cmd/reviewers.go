package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/mhristof/gi/git"
	"github.com/mhristof/gi/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var reviewersCmd = &cobra.Command{
	Use:   "reviewers",
	Short: "Find out who need to review git code.",
	Long: fmt.Sprintf(heredoc.Doc(`
		Find out people with code changes for files and repositories.

		If a file is passed, then 'git blame' is used as well as any merges
		that touch the file provided.

		If no argument is provided, then all files are checked from the repository

		Cache file: %s
	`), git.CacheLocation()),
	Run: func(cmd *cobra.Command, args []string) {
		branch, err := cmd.Flags().GetBool("branch")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panic("cannot get 'branch' flag")
		}

		log.WithFields(log.Fields{
			"git.Branch()": git.Branch(),
			"git.Main()":   git.Main(),
			"branch":       branch,
		}).Debug("current branch")

		if branch {
			args = util.Eval(fmt.Sprintf("git diff --name-only %s", git.Main()))
			// empty line at the end of the array
			args = args[0 : len(args)-1]
		}

		if len(args) == 0 {
			args = git.Files()
		}

		g := git.NewFromFiles(args)

		log.WithFields(log.Fields{
			"g": fmt.Sprintf("%+v", g),
		}).Debug("Reviewers")

		fmt.Println(strings.Join(g.Reviewers(), ","))
	},
}

func init() {
	rootCmd.AddCommand(reviewersCmd)
}
