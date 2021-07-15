package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3 sync",
	Short: "Sync to AWS s3",
	Long:  `Use gzsync s3 sync command for full information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use: gzsync s3 sync")
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
}
