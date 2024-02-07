/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"silabs/get-zap/jf"

	"github.com/spf13/cobra"
)

var rtDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads a file from Artifactory",
	Long:  `Performs a download of one or more files from Artifactory.`,
	Run: func(cmd *cobra.Command, args []string) {
		jf.ArtifactoryDownload(ReadArtifactoryConfiguration())
	},
}

func init() {
	rtCmd.AddCommand(rtDownloadCmd)
}
