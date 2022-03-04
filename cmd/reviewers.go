package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/mhristof/gi/git"
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
		reviewers := map[string]int

		for _, file := range args {
			rev, err := gg.Reviewers(file)
			if err != nil {
				log.WithFields(log.Fields{
					"err":  err,
					"file": file,
				}).Error("cannot get reviewers for file")

				continue
			}
		}

		fmt.Println(fmt.Sprintf("reviewers: %+v %T", reviewers, reviewers))
	},
}

func init() {
	rootCmd.AddCommand(reviewersCmd)
}
