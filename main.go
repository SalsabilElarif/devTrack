package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Task represents a single task with description, timestamp, and done status
type Task struct {
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	Done        bool      `json:"done"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: devTrack [add <task>] | [report] | [done <task #>] | [delete <task #>]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		task := os.Args[2]
		for i := 3; i < len(os.Args); i++ {
			task += " " + os.Args[i]
		}
		addTask(task)
	case "report":
		showTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task number to mark done.")
			return
		}
		markTaskDone(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task number to delete.")
			return
		}
		deleteTask(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
	}
}

// addTask adds a new task with current timestamp and saves it
func addTask(description string) {
	tasks := loadTasks()
	tasks = append(tasks, Task{
		Description: description,
		Timestamp:   time.Now(),
		Done:        false,
	})
	saveTasks(tasks)
	fmt.Println("Added task:", description)
}

// showTasks loads and displays all tasks with their status and timestamps
func showTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Your tasks:")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s %s\n", i+1, task.Timestamp.Format("2006-01-02 15:04"), status, task.Description)
	}
}

// markTaskDone marks the specified task as done and saves the list
func markTaskDone(numStr string) {
	tasks := loadTasks()

	num, err := strconv.Atoi(numStr)
	if err != nil || num < 1 || num > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks[num-1].Done = true
	saveTasks(tasks)
	fmt.Printf("Marked task #%d as done.\n", num)
}

// deleteTask removes the specified task from the list and saves it
func deleteTask(numStr string) {
	tasks := loadTasks()

	num, err := strconv.Atoi(numStr)
	if err != nil || num < 1 || num > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	task := tasks[num-1]
	tasks = append(tasks[:num-1], tasks[num:]...)
	saveTasks(tasks)
	fmt.Printf("Deleted task #%d: %s\n", num, task.Description)
}

// loadTasks reads tasks from tasks.json, returns empty slice if none found
func loadTasks() []Task {
	var tasks []Task
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return tasks
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error reading tasks:", err)
	}
	return tasks
}

// saveTasks saves the tasks slice to tasks.json file
func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
	}
}
