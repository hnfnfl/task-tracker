package cmd

import (
	"encoding/json"
	"io"
	"os"
)

const (
	tasksFile = "cmd/tasks.json"
)

// read task from tasks.json
func readTasks() ([]Task, error) {
	file, err := os.Open(tasksFile)
	if os.IsNotExist(err) {
		file, err = os.Create(tasksFile)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		return []Task{}, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		if err == io.EOF {
			return []Task{}, nil
		}

		return nil, err
	}

	return tasks, nil
}

// write task to tasks.json
func writeTasks(task Task) error {
	existingTasks, err := readTasks()
	if err != nil {
		return err
	}

	existingTasks = append(existingTasks, task)

	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(existingTasks)
	if err != nil {
		return err
	}

	return nil
}
