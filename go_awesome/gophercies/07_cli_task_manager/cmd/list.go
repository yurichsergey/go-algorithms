package cmd

import (
	"07_cli_task_manager/db"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task, err := db.ListTasks()
		if err != nil {
			fmt.Println("error listing tasks:", err)
			return
		}
		if len(task) == 0 {
			fmt.Println("no tasks found")
			return
		}

		fmt.Printf("You have the following tasks:\n")
		for _, t := range task {
			fmt.Printf("- %s\n", t)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
