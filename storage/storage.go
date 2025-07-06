package storage

import (
	"encoding/json"
	"os"
	"strconv"
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
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//handle if the file is empty adding both []
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() == 0 {
		// File is empty, start with opening bracket
		file.WriteString("[\n")
		//write the task as a json object (first entry, no comma)
		file.WriteString(`  {"ID":"` + id + `","TaskName":"` + taskName + `","Status":` + strconv.FormatBool(status) + `}` + "\n")
		file.WriteString("]")
	} else {
		// File has content, need to insert before the closing bracket
		// Read current content
		file.Seek(0, 0)
		content := make([]byte, info.Size())
		file.Read(content)

		// Remove the closing bracket and add comma + new entry + closing bracket
		contentStr := string(content)
		if len(contentStr) > 0 && contentStr[len(contentStr)-1] == ']' {
			// Remove the last character (closing bracket)
			contentStr = contentStr[:len(contentStr)-1]

			// Truncate file and write updated content
			file.Truncate(0)
			file.Seek(0, 0)
			file.WriteString(contentStr)
			file.WriteString(`,` + "\n")
			file.WriteString(`  {"ID":"` + id + `","TaskName":"` + taskName + `","Status":` + strconv.FormatBool(status) + `}` + "\n")
			file.WriteString("]")
		}
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
