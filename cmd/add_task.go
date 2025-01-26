package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func AddTask() {
	var (
		description string
		timestamp   = time.Now().Format("2006-01-02 15:04:05")
	)

	// set up flag
	addTaskCmd := flag.NewFlagSet("Add a new task", flag.ExitOnError)
	addTaskCmd.StringVar(&description, "desc", "", "Description of the task")

	// define usage
	addTaskCmd.Usage = func() {
		fmt.Println("Usage: task-tracker add [options]")
		addTaskCmd.PrintDefaults()
	}

	// check if the command line arguments are less than 3
	if len(os.Args) < 3 {
		addTaskCmd.Usage()
		return
	}

	// parse the command line arguments
	if err := addTaskCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing command line arguments:", err)
		return
	}

	// read from tasks file
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	// set the ID of the new task object
	newID := 1
	for _, task := range tasks {
		if task.ID >= newID {
			newID = task.ID + 1
		}
	}

	// create a new task object
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      StatusEnum(Todo),
		CreatedAt:   timestamp,
		UpdatedAt:   timestamp,
	}

	// write to task.json
	if err := upsertTask(newTask); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Task added with ID:", newTask.ID)
}
