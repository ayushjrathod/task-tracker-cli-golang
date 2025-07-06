package task

import (
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
		println("No tasks found.")
		return
	}

	for _, task := range tasks {
		status := "Not Done"
		if task.Status {
			status = "Done"
		}
		println("ID:", task.ID, "Task Name:", task.TaskName, "Status:", status)
	}
}
