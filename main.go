package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task-manager/cmd"
)

func main() {
	// If arguments are provided, run as single command
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	// If no arguments, enter interactive mode
	fmt.Println("Welcome to Task Manager CLI")
	fmt.Println("Type 'exit' to quit the program")
	fmt.Println("Available commands: add, list, remove, done, clear, help")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("tm> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Check for exit command
		if input == "exit" || input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		// Skip empty input
		if input == "" {
			continue
		}

		// Parse the input into arguments like os.Args
		args := strings.Fields(input)

		// Execute the command with the parsed arguments
		cmd.ExecuteInteractive(args)
	}
}
