package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
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

// upsert task in tasks.json
func upsertTask(task Task) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	// find and update the task if it exists
	found := false
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			found = true
			break
		}
	}

	// if the task was not found, append it as a new task
	if !found {
		tasks = append(tasks, task)
	}

	// write tasks back to file
	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

// delete task from tasks.json
func deleteTask(id int, all bool) error {
	if all {
		// truncate the file
		file, err := os.Create(tasksFile)
		if err != nil {
			return err
		}
		defer file.Close()

		return nil
	}

	tasks, err := readTasks()
	if err != nil {
		return err
	}

	// delete the task if it exists
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	// write tasks back to file
	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

// print task
func printTasks(tasks []Task) {
	var (
		status string
	)

	// Create a new tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	// Print the header
	fmt.Fprintln(w, "ID\tDescription\tStatus\tCreated At\tUpdated At\t")
	fmt.Fprintln(w, "--\t-----------\t------\t----------\t----------\t")

	for _, t := range tasks {
		switch t.Status {
		case Todo:
			status = "Todo"
		case InProgress:
			status = "In Progress"
		case Done:
			status = "Done"
		}

		// Print each task in a tab-separated format
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t\n", t.ID, t.Description, status, t.CreatedAt, t.UpdatedAt)
	}

	w.Flush()
}

func findTaskByID(tasks []Task, id int) (Task, bool) {
	for _, task := range tasks {
		if task.ID == id {
			return task, true
		}
	}
	return Task{}, false
}
