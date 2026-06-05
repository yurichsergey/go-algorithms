package cmd

import (
	"07_cli_task_manager/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to a task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.AddTask(task)
		if err != nil {
			fmt.Printf("Error adding task %s: %v\n", task, err)
			return
		}
		fmt.Printf("Task added: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
