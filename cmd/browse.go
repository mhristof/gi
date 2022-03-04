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
		cwd, err := cmd.Flags().GetString("cwd")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panic("cannot retrieve cwd flag")
		}

		gg, err := git.New(cwd)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("cannot create git")
		}

		targets := []string{gg.Dir}

		if len(args) > 0 {
			targets = args
		}

		for _, target := range targets {
			if _, err := os.Stat(target); os.IsNotExist(err) {
				log.WithFields(log.Fields{
					"err":    err,
					"target": target,
				}).Error("not found")
			}

			fmt.Println(fmt.Sprintf("target: %+v %T", target, target))
		}

		line, err := cmd.Flags().GetInt("line")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panic("Cannot retrieve line arg")
		}

		fmt.Println(fmt.Sprintf("line: %+v %T", line, line))

		//url, err := repo.URL(args[0], line)
		//if err != nil {
		//log.WithFields(log.Fields{
		//"err": err,
		//}).Fatal("Cannot calculate url")
		//}
		//fmt.Println(url)
		// do something
	},
}

func init() {
	browseCmd.Flags().IntP("line", "l", -1, "Line number")
	rootCmd.AddCommand(browseCmd)
}
