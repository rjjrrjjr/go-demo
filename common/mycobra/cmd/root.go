package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jin",
	Short: "first cobra",
	Long:  "jin's first cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do jin", cmd.Name(), args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
