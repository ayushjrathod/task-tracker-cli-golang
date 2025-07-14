package cmd

import (
	"fmt"
	"strconv"

	"task-manager/task"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark a task as completed",
	Long:  `Mark a task as completed by specifying its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid number.\n", args[0])
			return
		}
		err = task.MarkTaskDone(taskId)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Task %d marked as done.\n", taskId)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
