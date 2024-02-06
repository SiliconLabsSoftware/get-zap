/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"silabs/get-zap/jf"

	"github.com/spf13/cobra"
)

// artifactoryCmd represents the artifactory command
var artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Tests artifactory API.",
	Long:  `Use this function to quickly test artifactory credentials and connectivity.`,
	Run: func(cmd *cobra.Command, args []string) {
		jf.ArtifactoryDownload(ReadArtifactoryConfiguration())
	},
}

func init() {
	rootCmd.AddCommand(artifactoryCmd)
}
