package cmd

import (
	"task-manager/task"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Display all tasks with their completion status.`,
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
