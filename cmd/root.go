package cmd

import (
	"fmt"
	"os"

	"github.com/mhristof/gi/git"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "devel"
	gg      *git.Repo
	cfgFile = ""
)

var rootCmd = &cobra.Command{
	Use:     "gi",
	Short:   "A git helper",
	Long:    `TODO: changeme`,
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		Verbose(cmd)

		cwd, err := cmd.Flags().GetString("cwd")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panic("cannot retrieve cwd flag")
		}

		gg, err = git.New(cwd)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("cannot create git")
		}
	},
}

// Verbose Increase verbosity.
func Verbose(cmd *cobra.Command) {
	verbose, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		log.Panic(err)
	}

	switch verbose {
	case 1:
		log.SetLevel(log.DebugLevel)
	case 2:
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("cwd", "C", ".", "Run as if git was started in <path> instead of the current working directory.")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase verbosity")
	rootCmd.PersistentFlags().BoolP("dryrun", "n", false, "Dry run")
}

// Execute The main function for the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gi")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warning("cannot read config")
	}
}
