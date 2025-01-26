package cmd

import (
	"flag"
	"fmt"
	"os"
)

func ListTasks() {
	var (
		todo       bool
		inProgress bool
		done       bool
	)

	// set up flag
	listTaskCmd := flag.NewFlagSet("List all tasks", flag.ExitOnError)
	listTaskCmd.BoolVar(&todo, "todo", false, "List all tasks that are todo")
	listTaskCmd.BoolVar(&inProgress, "in-progress", false, "List all tasks that are in progress")
	listTaskCmd.BoolVar(&done, "done", false, "List all tasks that are done")

	// define usage
	listTaskCmd.Usage = func() {
		fmt.Println("Usage: task-tracker list [options]")
		fmt.Println("do not use any options to list all tasks")
		listTaskCmd.PrintDefaults()
	}

	// check if the command line arguments are less than 2
	if len(os.Args) < 2 {
		listTaskCmd.Usage()
		return
	}

	// parse the command line arguments
	if err := listTaskCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing command line arguments:", err)
		return
	}

	// read from tasks file
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	// list tasks based on status or list all tasks
	var outputs []Task
	for _, task := range tasks {
		if (!todo && !inProgress && !done) ||
			(todo && task.Status == Todo) ||
			(inProgress && task.Status == InProgress) ||
			(done && task.Status == Done) {
			outputs = append(outputs, task)
		}
	}

	if len(outputs) == 0 {
		fmt.Println("No tasks found")
	} else {
		printTasks(outputs)
	}
}
