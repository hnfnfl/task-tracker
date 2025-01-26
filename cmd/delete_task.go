package cmd

import (
	"flag"
	"fmt"
	"os"
)

func DeleteTask() {
	var (
		id  int
		all bool
	)

	// set up flag
	deleteTaskCmd := flag.NewFlagSet("Delete an existing task", flag.ExitOnError)
	deleteTaskCmd.IntVar(&id, "id", 0, "ID of the task to delete")
	deleteTaskCmd.BoolVar(&all, "all", false, "Delete all tasks")

	// define usage
	deleteTaskCmd.Usage = func() {
		fmt.Println("Usage: task-tracker delete [options]")
		deleteTaskCmd.PrintDefaults()
	}

	// check if the command line arguments are less than 3
	if len(os.Args) < 3 {
		deleteTaskCmd.Usage()
		return
	}

	// parse the command line arguments
	if err := deleteTaskCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing command line arguments:", err)
		return
	}

	// check if all flag is set
	if all {
		fmt.Println("You are about to delete all tasks. Are you sure? (y/n): ")
		var confirm string
		fmt.Scan(&confirm)
		if confirm != "y" {
			fmt.Println("Aborted")
			return
		}

		// write to task.json
		if err := deleteTask(0, true); err != nil {
			fmt.Println("Error deleting all tasks:", err)
			return
		}

		fmt.Println("All tasks deleted successfully")
		return
	}

	// check if ID is 0
	if id == 0 {
		fmt.Println("ID cannot be 0")
		return
	}

	// read from tasks file
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	// find the task with the given ID
	task, found := findTaskByID(tasks, id)
	if !found {
		fmt.Println("Task not found")
		return
	}

	// delete the task based on ID
	if err := deleteTask(task.ID, false); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Printf("Task with ID %d deleted successfully\n", id)
}
