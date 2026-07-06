package main

import (
	"fmt"
	"os"
)

const usage = `Usage: task-cli <command> [arguments]

Commands:
  add <description>               Add a new task
  update <id> <description>       Update a task description
  delete <id>                     Delete a task
  mark-in-progress <id>           Mark a task as in progress
  mark-done <id>                  Mark a task as done
  list [done|todo|in-progress]    List tasks
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(os.Args[2:])
	case "update":
		handleUpdate(os.Args[2:])
	case "delete":
		handleDelete(os.Args[2:])
	case "mark-in-progress":
		handleMark(os.Args[2:], StatusInProgress)
	case "mark-done":
		handleMark(os.Args[2:], StatusDone)
	case "list":
		handleList(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n%s", command, usage)
		os.Exit(1)
	}
}

func handleAdd(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error: missing description.\nUsage: task-cli add <description>")
		os.Exit(1)
	}

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	desc := args[0]
	tasks, task := AddTask(tasks, desc)

	if err := SaveTasks(tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func handleUpdate(args []string) {
	fmt.Println("update: not yet implemented")
}

func handleDelete(args []string) {
	fmt.Println("delete: not yet implemented")
}

func handleMark(args []string, status string) {
	fmt.Println("mark: not yet implemented")
}

func handleList(args []string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	status := ""
	if len(args) > 0 {
		status = args[0]
	}

	filtered := FilterTasks(tasks, status)
	if len(filtered) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, t := range filtered {
		PrintTask(t)
	}
}
