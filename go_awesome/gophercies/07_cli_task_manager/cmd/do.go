package cmd

import (
	"07_cli_task_manager/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You must specify a task name")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error parsing task ID: %v\n", err)
			return
		}
		err = db.CompleteTask(id)
		if err != nil {
			fmt.Printf("Error completing task: %v\n", err)
			return
		}
		fmt.Printf("Task marked as complete: %v\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
