package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrTaskNotFound = errors.New("task not found")

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

func findTaskIndex(tasks []Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func UpdateTask(tasks []Task, id int, description string) ([]Task, error) {
	i := findTaskIndex(tasks, id)
	if i == -1 {
		return tasks, ErrTaskNotFound
	}

	tasks[i].Description = description
	tasks[i].UpdatedAt = time.Now()
	return tasks, nil
}

func MarkTask(tasks []Task, id int, status string) ([]Task, error) {
	i := findTaskIndex(tasks, id)
	if i == -1 {
		return tasks, ErrTaskNotFound
	}

	tasks[i].Status = status
	tasks[i].UpdatedAt = time.Now()
	return tasks, nil
}

func DeleteTask(tasks []Task, id int) ([]Task, error) {
	i := findTaskIndex(tasks, id)
	if i == -1 {
		return tasks, ErrTaskNotFound
	}

	tasks = append(tasks[:i], tasks[i+1:]...)
	return tasks, nil
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
