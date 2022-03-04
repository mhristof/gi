package cmd

// var reviewersCmd = &cobra.Command{
// 	Use:   "reviewers",
// 	Short: "Find out who need to review git code.",
// 	Long: fmt.Sprintf(heredoc.Doc(`
// 		Find out people with code changes for files and repositories.

// 		If a file is passed, then 'git blame' is used as well as any merges
// 		that touch the file provided.

// 		If no argument is provided, then all files are checked from the repository

// 		Cache file: %s
// 	`), git.CacheLocation()),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		branch, err := cmd.Flags().GetBool("branch")
// 		if err != nil {
// 			log.WithFields(log.Fields{
// 				"err": err,
// 			}).Panic("cannot get 'branch' flag")
// 		}

// 		branchName, _ := git.Branch()

// 		log.WithFields(log.Fields{
// 			"git.Branch()": branchName,
// 			"git.Main()":   git.Main(),
// 			"branch":       branch,
// 			"args":         args,
// 		}).Debug("current branch")

// 		if len(args) == 0 {
// 			if branch {
// 				args = util.Eval(fmt.Sprintf("git diff --name-only %s", git.Main()))
// 				// empty line at the end of the array
// 				args = args[0 : len(args)-1]
// 			} else {
// 				args = git.Files()
// 			}
// 		}

// 		g := git.NewFromFiles(args)

// 		log.WithFields(log.Fields{
// 			"g": fmt.Sprintf("%+v", g),
// 		}).Debug("Reviewers")

// 		fmt.Println(strings.Join(g.Reviewers(), ","))
// 	},
// }

// func init() {
// 	branch := false
// 	branchName, err := git.Branch()
// 	if err == nil && branchName == git.Main() {
// 		branch = true
// 	}

// 	rootCmd.PersistentFlags().BoolP(
// 		"branch", "b", branch,
// 		"Detect reviewers for current branch. Enabled when branch name is not 'main'",
// 	)
// 	rootCmd.PersistentFlags().StringSliceP(
// 		"bots", "",
// 		[]string{"semantic-release-bot@martynus.net"},
// 		"Bot list definition. Used with --human",
// 	)
// 	rootCmd.PersistentFlags().BoolP("human", "H", true, "Show human reviewers only.")
// 	rootCmd.AddCommand(reviewersCmd)
// }
