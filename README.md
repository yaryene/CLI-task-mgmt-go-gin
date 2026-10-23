# Task API

A small command-line task manager written in Go. You can add tasks, list them, mark them as done and delete them from an interactive menu.

Tasks are kept in memory only, so they are lost when the program exits.

## Requirements

- Go 1.27.3 or newer (see `go.mod`)

## Run

```sh
go run .
```

Or build a binary:

```sh
go build -o task-api .
./task-api
```

## Usage

When the program starts, it shows this menu:

```
Task Manager
1. Add Task
2. List Tasks
3. Mark Task as Done
4. Delete Task
5. Exit
Choose an option:
```

| Option | What it does |
|---|---|
| 1 | Asks for a name and a priority (1–5), then adds the task. IDs start at 1 and go up. |
| 2 | Lists every task with its ID, name, status and priority. |
| 3 | Asks for a task ID and marks that task as done. |
| 4 | Asks for a task ID and deletes that task. |
| 5 | Exits the program. |

Example output from option 2:

```
Task 1: Write docs [pending] (Priority: 2)
Task 2: Fix prod bug [done] (Priority: 5)
```

## Code overview

Everything is in `main.go`:

| Name | Kind | Purpose |
|---|---|---|
| `Task` | struct | One task: `ID`, `Name`, `Done`, `Priority` |
| `Task.Describe()` | method | Returns a one-line description of the task |
| `Task.IsUrgent()` | method | `true` when priority is 3 or higher and the task is not done |
| `findTaskIndex` | function | Finds a task's position in the slice by ID, or returns an error |
| `addTask` | function | Appends a new task and returns the updated slice |
| `markTaskDone` | function | Sets `Done` to `true` for the task with the given ID |
| `deleteTask` | function | Removes the task with the given ID |
| `readLine` | function | Prints a prompt and reads one line of input |

## Known limitations

- Tasks are not saved to disk.
- A non-numeric ID or priority is read as `0`, and priority is not checked against the 1–5 range.
- If input ends (for example after Ctrl+D), the menu loops forever. Use option 5 to exit.
- Gin is listed in `go.mod`, but the program does not run an HTTP server yet.
