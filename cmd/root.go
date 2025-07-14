package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tm",
	Short: "A lightweight command-line task management tool",
	Long:  `Task Manager CLI - Manage your tasks with simple commands while maintaining data persistence through local JSON storage.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExecuteInteractive(args []string) {
	// Create a new command instance for interactive mode
	cmd := &cobra.Command{
		Use:   "tm",
		Short: "A lightweight command-line task management tool",
		Long:  `Task Manager CLI - Manage your tasks with simple commands while maintaining data persistence through local JSON storage.`,
	}

	// Add all subcommands
	cmd.AddCommand(addCmd)
	cmd.AddCommand(listCmd)
	cmd.AddCommand(removeCmd)
	cmd.AddCommand(doneCmd)
	cmd.AddCommand(clearCmd)

	// Set the arguments and execute
	cmd.SetArgs(args)
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

func init() {
	// Add global flags here if needed
}
