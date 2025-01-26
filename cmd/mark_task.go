package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func MarkTask() {
	var id int

	// set up flag
	markTaskCmd := flag.NewFlagSet("Mark a task as in progress or done", flag.ExitOnError)
	markTaskCmd.IntVar(&id, "id", 0, "ID of the task to mark")

	// define usage
	markTaskCmd.Usage = func() {
		fmt.Println("Usage: task-tracker mark-in-progress|mark-done [options]")
		markTaskCmd.PrintDefaults()
	}

	// check if the command line arguments are less than 3
	if len(os.Args) < 3 {
		markTaskCmd.Usage()
		return
	}

	// parse the command line arguments
	if err := markTaskCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing command line arguments:", err)
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

	// update task status
	if os.Args[1] == "mark-in-progress" {
		task.Status = StatusEnum(InProgress)
	} else {
		task.Status = StatusEnum(Done)
	}
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// write to task.json
	if err := upsertTask(task); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Task updated successfully")
}
