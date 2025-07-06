package task

import (
	"fmt"
	"os"
	"task-manager/storage"
	"task-manager/utils"
)

func NewTask(taskName string) storage.Task {
	//check if tasks.json exists, if not create it
	if !storage.FileExists("tasks.json") {
		storage.CreateFile("tasks.json")
	}

	//adding new task to tasks.json
	task := storage.Task{
		ID:       utils.GenerateID(),
		TaskName: taskName,
		Status:   false,
	}

	storage.AddTaskToFile(task.ID, task.TaskName, task.Status, "tasks.json")
	return task
}

func ListTasks() {
	tasks, err := storage.ReadTasksFromFile("tasks.json")
	if err != nil {
		panic(err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		status := "Not Done"
		if task.Status {
			status = "Done"
		}
		fmt.Printf("ID: %s | %s | Status: %s\n", task.ID, task.TaskName, status)
	}
}

// Remove task by ID
func RemoveTask(taskID int) error {
	if !storage.FileExists("tasks.json") {
		return fmt.Errorf("no tasks file found")
	}

	taskIDStr := fmt.Sprintf("%d", taskID)
	_, err := storage.RemoveTaskFromFile(taskIDStr, "tasks.json")
	if err != nil {
		if err == os.ErrNotExist {
			return fmt.Errorf("task with ID %d not found", taskID)
		}
		return err
	}

	return nil
}

// Mark task as done
func MarkTaskDone(taskID int) error {
	if !storage.FileExists("tasks.json") {
		return fmt.Errorf("no tasks file found")
	}

	taskIDStr := fmt.Sprintf("%d", taskID)

	_, err := storage.UpdateTaskStatus(taskIDStr, true, "tasks.json")
	if err != nil {
		return err
	}

	return nil
}

// Clear all tasks
func ClearAllTasks() error {
	if !storage.FileExists("tasks.json") {
		return fmt.Errorf("no tasks file found")
	}

	_, err := storage.ClearAllTasks("tasks.json")
	if err != nil {
		return err
	}

	return nil
}
