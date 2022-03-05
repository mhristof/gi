package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/mhristof/gi/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var reviewersCmd = &cobra.Command{
	Use:   "reviewers",
	Short: "Find out who need to review git code.",
	Long: heredoc.Doc(`
		Find out people with code changes for files and repositories.

		If a file is provided, only that file will be checked.
		If no files are provided, the reviewers for the current branch will
		be calculated.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		reviewers := map[string]int{}
		var err error

		if len(args) == 0 {
			reviewers, err = gg.BranchReviewers()
			if err != nil {
				panic(err)
			}
		} else {
			for _, file := range args {
				fileR, err := gg.Reviewers(file)
				if err != nil {
					log.WithFields(log.Fields{
						"err":  err,
						"file": fileR,
					}).Warning("cannot get reviewers")

					continue
				}

				reviewers = util.MapMerge(reviewers, fileR)
			}
		}

		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			panic(err)
		}

		author := []string{}
		ignore := viper.GetStringSlice("ignore")
		list := util.SortMap(reviewers)

		for i := 0; i < count; i++ {
			if i >= len(list) {
				break
			}

			found := false
			for _, ignored := range ignore {
				if ignored == list[i].Key {
					found = true
					break
				}
			}

			if found {
				continue
			}

			author = append(author, list[i].Key)
		}

		fmt.Println(strings.Join(author, ","))
	},
}

func init() {
	reviewersCmd.PersistentFlags().StringSliceP("ignore", "i", []string{}, "Ignore authors")
	reviewersCmd.PersistentFlags().IntP("count", "c", 4, "Number of reviewers to show")

	viper.BindPFlag("ignore", reviewersCmd.PersistentFlags().Lookup("ignore"))

	rootCmd.AddCommand(reviewersCmd)
}
