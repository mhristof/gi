package cmd

import (
	"fmt"
	"os"

	"github.com/mhristof/gi/git"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:     "browse",
	Aliases: []string{"b"},
	Short:   "Find out the URL for any git item",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.WithFields(log.Fields{
				"args[0]": args[0],
			}).Error("Does not exist")
		}

		repo, err := git.New(args[0])
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("Cant create a repo")
		}

		line, err := cmd.Flags().GetInt("line")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panic("Cannot retrieve line arg")

		}

		url, err := repo.URL(args[0], line)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("Cannot calculate url")

		}
		fmt.Println(url)
		// do something
	},
}

func init() {
	browseCmd.Flags().IntP("line", "l", -1, "Line number")
	rootCmd.AddCommand(browseCmd)
}
