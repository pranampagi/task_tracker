package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseID(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid task ID: %q", s)
	}
	return id, nil
}

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
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: missing arguments.\nUsage: task-cli update <id> <description>")
		os.Exit(1)
	}

	id, err := parseID(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	tasks, err = UpdateTask(tasks, id, args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := SaveTasks(tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task updated successfully")
}

func handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error: missing task ID.\nUsage: task-cli delete <id>")
		os.Exit(1)
	}

	id, err := parseID(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	tasks, err = DeleteTask(tasks, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := SaveTasks(tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task deleted successfully")
}

func handleMark(args []string, status string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: missing task ID.\nUsage: task-cli mark-%s <id>\n", status)
		os.Exit(1)
	}

	id, err := parseID(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	tasks, err = MarkTask(tasks, id, status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := SaveTasks(tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task marked as %s successfully\n", status)
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
