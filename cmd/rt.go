/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"silabs/get-zap/jf"

	"github.com/spf13/cobra"
)

// rtCmd represents the rt command
var rtCmd = &cobra.Command{
	Use:   "rt",
	Short: "Artifactory related commands",
	Long:  `These commands are used to interact with artifactory.`,
	Run: func(cmd *cobra.Command, args []string) {
		jf.ArtifactoryDownload(ReadArtifactoryConfiguration())
	},
}

func init() {
	rootCmd.AddCommand(rtCmd)
}
