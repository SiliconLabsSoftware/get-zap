/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "Github related commands",
	Long:  `These commands are used to interact with github.`,
}

func init() {
	rootCmd.AddCommand(ghCmd)
}
