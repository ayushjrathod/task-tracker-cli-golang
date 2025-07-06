package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"task-manager/task"
)

func main() {
	fmt.Println("Welcome to Task Manager CLI")
	fmt.Println("Type 'exit' to quit the program")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove the newline character

		var (
			commandType     string
			commandArgument string
		)

		// Parse command and argument
		spaceIndex := -1
		for i := 0; i < len(input); i++ {
			if input[i] == ' ' {
				spaceIndex = i
				break
			}
		}

		if spaceIndex != -1 {
			commandType = input[:spaceIndex]
			commandArgument = input[spaceIndex+1:]
		} else {
			commandType = input
		}

		// Check for exit command
		if commandType == "exit" {
			fmt.Println("Goodbye! Thanks for using Task Manager CLI")
			break
		}

		var (
			taskName string
			taskId   int
		)

		if commandArgument != "" {
			//use regex to find whether commandArgument is a task name or task id
			// For simplicity, let's assume if it starts with a digit, it's an ID,
			if commandArgument[0] >= '0' && commandArgument[0] <= '9' {
				// Convert commandArgument to int using strconv.Atoi
				var err error
				taskId, err = strconv.Atoi(commandArgument)
				if err != nil {
					fmt.Println("Invalid task ID:", commandArgument)
					continue
				}
			} else {
				taskName = commandArgument
			}
		}

		switch commandType {
		case "add":
			task.NewTask(taskName)
			fmt.Println("Task added:", taskName)
		case "list":
			fmt.Println("Listing all tasks...")
			task.ListTasks()
		case "remove":
			fmt.Println("Removing a task...", taskId)
		case "done":
			fmt.Println("Marking task as done...")
		case "undo":
			fmt.Println("Undoing the last action...")
		case "clear":
			fallthrough
		case "reset":
			fmt.Println("Clearing all tasks...")
		default:
			fmt.Println("Unknown command. Available commands: add, list, remove, done, undo, clear, reset or exit.")
		}
	}
}
