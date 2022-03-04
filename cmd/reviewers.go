package cmd

import (
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/mhristof/gi/git"
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
	Args: func(cmd *cobra.Command, args []string) error {
		branch, err := cmd.Flags().GetBool("branch")
		if err != nil {
			panic(err)
		}

		if len(args) > 0 && branch {
			return errors.New("cannot use branch flag and provide args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//branch, err := cmd.Flags().GetBool("branch")
		//if err != nil {
		//panic(err)
		//}

		// if branch {
		reviewers, err := gg.BranchReviewers()
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("reviewers: %+v %T", reviewers, reviewers))
		//}
	},
}

func init() {
	reviewersCmd.PersistentFlags().BoolP("branch", "b", true, "Calculate reviewers for the current branch changes")
	rootCmd.AddCommand(reviewersCmd)
}
