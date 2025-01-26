package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func UpdateTask() {
	var (
		id          int
		description string
	)

	// set up flag
	updateTaskCmd := flag.NewFlagSet("Update an existing task", flag.ExitOnError)
	updateTaskCmd.IntVar(&id, "id", 0, "ID of the task to update")
	updateTaskCmd.StringVar(&description, "desc", "", "Description of the task")

	// define usage
	updateTaskCmd.Usage = func() {
		fmt.Println("Usage: task-tracker update [options]")
		updateTaskCmd.PrintDefaults()
	}

	// check if the command line arguments are less than 3
	if len(os.Args) < 3 {
		updateTaskCmd.Usage()
		return
	}

	// parse the command line arguments
	if err := updateTaskCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing command line arguments:", err)
		return
	}

	// check if ID is 0
	if id == 0 {
		fmt.Println("ID cannot be 0")
		return
	}

	// check if description is empty
	if description == "" {
		fmt.Println("Description empty, nothing to update")
		return
	}

	// read from tasks file
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	// find the task by ID
	task, found := findTaskByID(tasks, id)
	if !found {
		fmt.Println("Task ID not found")
		return
	}

	// update the task
	task.Description = description
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// write to task.json
	if err := upsertTask(task); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Task updated successfully")
}
