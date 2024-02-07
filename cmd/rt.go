/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rtCmd = &cobra.Command{
	Use:   "rt",
	Short: "Artifactory related commands",
	Long:  `These commands are used to interact with artifactory.`,
}

func init() {
	rootCmd.AddCommand(rtCmd)
}
