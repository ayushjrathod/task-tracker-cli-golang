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
			if taskName == "" {
				fmt.Println("Error: Please provide a task name. Usage: add <task_name>")
				continue
			}
			task.NewTask(taskName)
		case "list":
			task.ListTasks()
		case "remove":
			if taskId == 0 {
				fmt.Println("Error: Please provide a task ID. Usage: remove <task_id>")
				continue
			}
			err := task.RemoveTask(taskId)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "done":
			if taskId == 0 {
				fmt.Println("Error: Please provide a task ID. Usage: done <task_id>")
				continue
			}
			err := task.MarkTaskDone(taskId)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "clear":
			fallthrough
		case "reset":
			err := task.ClearAllTasks()
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			fmt.Println("Unknown command. Available commands: add, list, remove, done, clear, reset or exit.")
		}
	}
}
