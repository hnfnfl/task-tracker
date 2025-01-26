package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
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

// print task
func printTasks(tasks []Task) {
	var (
		sb     strings.Builder
		status string
	)

	for _, t := range tasks {
		switch t.Status {
		case Todo:
			status = "Todo"
		case InProgress:
			status = "In Progress"
		case Done:
			status = "Done"
		}

		sb.WriteString("--------------------------------\n")
		sb.WriteString(fmt.Sprintf("ID: %d\n", t.ID))
		sb.WriteString(fmt.Sprintf("Description: %s\n", t.Description))
		sb.WriteString(fmt.Sprintf("Status: %s\n", status))
		sb.WriteString(fmt.Sprintf("Created At: %s\n", t.CreatedAt))
		sb.WriteString(fmt.Sprintf("Updated At: %s\n", t.UpdatedAt))
	}

	fmt.Println(sb.String())
}
