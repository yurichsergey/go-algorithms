package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is a fake \"do\" command, args: %v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
