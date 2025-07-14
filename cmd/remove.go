package cmd

import (
	"fmt"
	"strconv"

	"task-manager/task"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [task ID]",
	Short: "Remove a task",
	Long:  `Remove a task from your task list by specifying its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid number.\n", args[0])
			return
		}
		err = task.RemoveTask(taskId)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Task %d removed successfully.\n", taskId)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
