/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"silabs/get-zap/github"

	"github.com/spf13/cobra"
)

// selfCheckCmd represents the selfCheck command
var selfCheckCmd = &cobra.Command{
	Use:   "selfcheck",
	Short: "Runs a simple self-check.",
	Long:  "Execute a self-check of connectivity.",
	Run: func(cmd *cobra.Command, args []string) {
		github.SelfCheck(ReadGithubConfiguration())
	},
}

func init() {
	rootCmd.AddCommand(selfCheckCmd)
}
