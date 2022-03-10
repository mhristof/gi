package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mhristof/gi/jira"
	"github.com/mhristof/gi/keychain"
	"github.com/mhristof/gi/util"
	"github.com/spf13/cobra"
)

var featCmd = &cobra.Command{
	Use:     "feat",
	Aliases: []string{"f"},
	Short:   "Create git feature branches",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, "-")

		if len(args) == 1 {
			jiras, _ := UserGet(cmd, args, args[0])

			for _, jira := range jiras {
				if strings.HasPrefix(jira, args[0]) {
					title = jira

					break
				}
			}
		}

		title = strings.ReplaceAll(title, " ", "-")
		title = strings.ReplaceAll(title, "\t", "-")
		title = strings.ReplaceAll(title, ":", "-")
		title = strings.ReplaceAll(title, "--", "-")
		title = strings.ReplaceAll(title, ",", "-")

		util.Bash(fmt.Sprintf("git checkout -b %s", title))
	},
	ValidArgsFunction: UserGet,
}

func UserGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	jiraToken, err := cmd.Flags().GetString("jira-token")
	if err != nil {
		panic(err)
	}

	jiraURL, err := cmd.Flags().GetString("jira-url")
	if err != nil {
		panic(err)
	}

	jiraUser, err := cmd.Flags().GetString("jira-user")
	if err != nil {
		panic(err)
	}

	clearCache, err := cmd.Flags().GetBool("clear-cache")
	if err != nil {
		panic(err)
	}

	j := jira.New(jiraURL, jiraUser, jiraToken)

	if clearCache {
		j.ClearCache()
	}

	issues := j.Issues()

	ret := make([]string, len(issues))
	for i, issue := range issues {
		ret[i] = strings.Join([]string{issue.Key, issue.Fields.Summary}, "\t")
	}

	return ret, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	token := os.Getenv("JIRA_TOKEN")
	if token == "" {
		token = keychain.Item("JIRA_TOKEN")
	}
	url := os.Getenv("JIRA_URL")
	user := os.Getenv("JIRA_USER")

	featCmd.PersistentFlags().BoolP("clear-cache", "c", false, "Clear cache")
	featCmd.PersistentFlags().StringP("jira-token", "t", token, "Jira token")
	featCmd.PersistentFlags().StringP("jira-url", "", url, "Jira url")
	featCmd.PersistentFlags().StringP("jira-user", "u", user, "jira username")
	rootCmd.AddCommand(featCmd)
}
