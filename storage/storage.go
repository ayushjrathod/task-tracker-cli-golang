package storage

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID       string `json:"ID"`
	TaskName string `json:"TaskName"`
	Status   bool   `json:"Status"`
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// create a json file
func CreateFile(filename string) string {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return filename
}

// add task to json file
func AddTaskToFile(id, taskName string, status bool, filename string) {
	// Read existing tasks
	var tasks []Task
	if FileExists(filename) {
		existingTasks, err := ReadTasksFromFile(filename)
		if err != nil {
			// If file exists but can't be read, start with empty slice
			tasks = []Task{}
		} else {
			tasks = existingTasks
		}
	}

	// Add new task
	newTask := Task{
		ID:       id,
		TaskName: taskName,
		Status:   status,
	}
	tasks = append(tasks, newTask)

	// Write all tasks back to file
	err := WriteTasksToFile(tasks, filename)
	if err != nil {
		panic(err)
	}
}

// Read tasks from json file
func ReadTasksFromFile(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Write tasks to json file (overwrites existing content)
func WriteTasksToFile(tasks []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

// Remove task by ID
func RemoveTaskFromFile(taskID string, filename string) (Task, error) {
	tasks, err := ReadTasksFromFile(filename)
	if err != nil {
		return Task{}, err
	}

	var removedTask Task
	var updatedTasks []Task
	found := false

	for _, task := range tasks {
		if task.ID == taskID {
			removedTask = task
			found = true
		} else {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if !found {
		return Task{}, os.ErrNotExist
	}

	err = WriteTasksToFile(updatedTasks, filename)
	return removedTask, err
}

// Update task status by ID
func UpdateTaskStatus(taskID string, status bool, filename string) (Task, error) {
	tasks, err := ReadTasksFromFile(filename)
	if err != nil {
		return Task{}, err
	}

	var updatedTask Task
	found := false

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Status = status
			updatedTask = tasks[i]
			found = true
			break
		}
	}

	if !found {
		return Task{}, os.ErrNotExist
	}

	err = WriteTasksToFile(tasks, filename)
	return updatedTask, err
}

// Clear all tasks
func ClearAllTasks(filename string) ([]Task, error) {
	tasks, err := ReadTasksFromFile(filename)
	if err != nil {
		return nil, err
	}

	emptyTasks := []Task{}
	err = WriteTasksToFile(emptyTasks, filename)
	return tasks, err
}
