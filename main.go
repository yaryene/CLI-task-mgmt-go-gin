package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Task struct {
	ID int
	Name string
	Done bool
	Priority int
}
// method to describe a task
// The (t Task) before the name is the receiver. It makes Describe belong to Task, so you call it on a task value:
func (t Task) Describe() string { // receiver: this attaches the function to the Task type
	status := "pending"
	if t.Done {
		status = "done"
	}
	return fmt.Sprintf("Task %d: %s [%s] (Priority: %d)", t.ID, t.Name, status, t.Priority)
}

// method to check if a task is urgent
func (t Task) IsUrgent() bool {
	return t.Priority >= 3 && !t.Done
}

// function to find a task by ID in a slice of tasks
func findTaskIndex(tasks []Task, id int) (int, error) {
	for i, t := range tasks {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("task with ID %d not found", id)
}

func addTask(tasks []Task, nextID int, name string, priority int) []Task{
	newTask := Task{ID: nextID, Name: name, Done: false, Priority: priority}
	tasks = append(tasks, newTask)
	return tasks
}
func markTaskDone(tasks []Task, id int) ([]Task, error) {
	index, err := findTaskIndex(tasks, id)
	if err != nil {
		return tasks, err
	}
	tasks[index].Done = true
	return tasks, nil
}
func deleteTask(tasks []Task, id int) ([]Task, error) {
	index, err := findTaskIndex(tasks, id)
	if err != nil {
		return tasks, err
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	return tasks, nil
}
func readLine(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {

	var tasks []Task
	nextID := 1

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		choice := readLine("Choose an option: ")

		switch choice {
		case "1":
			name := readLine("Enter task name: ")
			priorityStr := readLine("Enter task priority (1-5): ")
			priority, _ := strconv.Atoi(priorityStr)
			tasks = addTask(tasks, nextID, name, priority)
			nextID++
			fmt.Println("Task added.")
		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				for _, t := range tasks {
					fmt.Println(t.Describe())
				}
			}
		case "3":
			idStr := readLine("Enter task ID to mark as done: ")
			id, _ := strconv.Atoi(idStr)
			var err error
			tasks, err = markTaskDone(tasks, id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task marked as done.")
			}
		case "4":
			idStr := readLine("Enter task ID to delete: ")
			id, _ := strconv.Atoi(idStr)
			var err error
			tasks, err = deleteTask(tasks, id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted.")
			}
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
