package cmd

import (
	"fmt"
	"strings"

	"task-manager/task"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Long:  `Add a new task to your task list with the specified name.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		task.NewTask(taskName)
		fmt.Printf("Task added: %s\n", taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
