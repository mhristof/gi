package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:     "browse",
	Aliases: []string{"b"},
	Short:   "Find out the URL for any git item",
	Run: func(cmd *cobra.Command, args []string) {
		targets := []string{gg.Dir}

		if len(args) > 0 {
			targets = args
		}

		line, err := cmd.Flags().GetInt("line")
		if err != nil {
			panic(err)
		}

		for _, target := range targets {
			realPath := filepath.Join(gg.Dir, target)
			if _, err := os.Stat(realPath); os.IsNotExist(err) {
				log.WithFields(log.Fields{
					"err":    err,
					"target": target,
				}).Error("not found")
			}

			url, err := gg.WebURL(target, line)
			if err != nil {
				log.WithFields(log.Fields{
					"err":  err,
					"path": target,
				}).Error("cannot calculate URL")
			}

			fmt.Print(url, "\n")
		}
	},
}

func init() {
	browseCmd.Flags().IntP("line", "l", -1, "Line number")
	rootCmd.AddCommand(browseCmd)
}
