/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"silabs/get-zap/gh"

	"github.com/spf13/cobra"
)

// selfCheckCmd represents the selfCheck command
var selfCheckCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists releases or release artifacts for a specific release.",
	Long: `This command will list the specified information from github.

Without any additional arguments, it will print all available releases for a given repo.
When specified with the --release tag, it will print the available assets for that release.`,
	Run: func(cmd *cobra.Command, args []string) {
		gh.ListGithub(ReadGithubConfiguration())
	},
}

func init() {
	rootCmd.AddCommand(selfCheckCmd)
}
