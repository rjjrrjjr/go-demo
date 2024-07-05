package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "short version",
	Long:  "long version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do version", cmd.Name(), args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
