package main

import (
	"flag"
	"fmt"
	"os"
	"task-tracker/cmd"
)

var showUsage = func() {
	fmt.Println("Usage: task-tracker [command] [options]")
	fmt.Println(`Commands:
     add: Add a new task to the tracker
     update: Update an existing task in the tracker
     delete: Delete a task from the tracker
     list: List all tasks in the tracker
     mark-in-progress: Mark a task as in progress
     mark-done: Mark a task as done`)
}

func main() {
	flag.Usage = showUsage

	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		cmd.AddTask()
	case "update":
		cmd.UpdateTask()
	case "delete":
		cmd.DeleteTask()
	case "list":
		cmd.ListTasks()
	case "mark-in-progress", "mark-done":
		// TODO: Implement mark-in-progress and mark-done commands
		// cmd.MarkTask()
	default:
		flag.Usage()
	}
}
