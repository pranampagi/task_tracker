package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(tasks []Task, description string) ([]Task, Task) {
	now := time.Now()
	task := Task{
		ID:          nextID(tasks),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	return tasks, task
}

func FilterTasks(tasks []Task, status string) []Task {
	if status == "" {
		return tasks
	}

	result := []Task{}
	for _, t := range tasks {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result
}

func statusLabel(status string) string {
	switch status {
	case StatusTodo:
		return "[todo]"
	case StatusInProgress:
		return "[in-progress]"
	case StatusDone:
		return "[done]"
	default:
		return "[unknown]"
	}
}

func PrintTask(t Task) {
	fmt.Printf("%d. %s %s\n", t.ID, statusLabel(t.Status), t.Description)
}
