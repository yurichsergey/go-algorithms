package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is a fake \"list\" command\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
