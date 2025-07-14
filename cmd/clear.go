package cmd

import (
	"fmt"

	"task-manager/task"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:     "clear",
	Short:   "Clear all tasks",
	Long:    `Remove all tasks from your task list.`,
	Aliases: []string{"reset"},
	Run: func(cmd *cobra.Command, args []string) {
		err := task.ClearAllTasks()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("All tasks cleared successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
