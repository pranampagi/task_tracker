# Task Tracker CLI

A simple command-line task tracker written in Go. Manage your to-do list from the terminal — add, update, delete, and mark tasks as they progress.

## Usage

```bash
# Build
go build -o task-cli .

# Add a task
task-cli add "Buy groceries"

# Update a task
task-cli update 1 "Buy groceries and cook dinner"

# Delete a task
task-cli delete 1

# Mark a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# List all tasks
task-cli list

# List tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Task Properties

Each task stored in `tasks.json` contains:

| Property      | Description                        |
| ------------- | ---------------------------------- |
| `id`          | Unique auto-increment identifier   |
| `description` | Short description of the task      |
| `status`      | Current status (`todo`, `in-progress`, or `done`) |
| `createdAt`   | ISO 8601 timestamp of creation     |
| `updatedAt`   | ISO 8601 timestamp of last update  |

## Build

Requires Go 1.21+. No external dependencies.

```bash
git clone https://github.com/pranampagi/task_tracker
cd task_tracker
go build -o task-cli .
./task-cli list
```

## Project Structure

```
task_tracker/
├── main.go       CLI entry point and command handlers
├── task.go       Task struct and business logic
├── storage.go    JSON file read/write
└── tasks.json    Auto-created task storage file
```
