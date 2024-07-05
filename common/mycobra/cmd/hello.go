package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "short hello",
	Long:  "long hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do hello", cmd.Name(), args)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
