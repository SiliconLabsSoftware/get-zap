/*
Copyright © 2024 Silicon Labs
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rtUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a file to artifactory",
	Long:  `Performs an upload of specified files to artifactory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
	},
}

func init() {
	rtCmd.AddCommand(rtUploadCmd)
}
