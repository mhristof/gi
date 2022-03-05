package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/mhristof/gi/git"
	"github.com/mhristof/gi/util"
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
		reviewers, err := gg.BranchReviewers()
		if err != nil {
			panic(err)
		}

		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			panic(err)
		}

		author := []string{}

		list := util.SortMap(reviewers)
		for i := 0; i < count; i++ {
			if i >= len(list) {
				break
			}

			author = append(author, list[i].Key)
		}

		fmt.Println(strings.Join(author, ","))
	},
}

func init() {
	reviewersCmd.PersistentFlags().IntP("count", "c", 4, "Number of reviewers to show")
	reviewersCmd.PersistentFlags().BoolP("branch", "b", true, "Calculate reviewers for the current branch changes")
	rootCmd.AddCommand(reviewersCmd)
}
