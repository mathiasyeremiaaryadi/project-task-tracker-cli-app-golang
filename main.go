package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func main() {
	tasks := InitialReadTask()

	if len(os.Args) < 1 {
		fmt.Println("Please provide an operation: add, update, delete, or list.")
		return
	}

	operation := os.Args[1]
	switch operation {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}

		id := os.Args[2]
		description := strings.Join(os.Args[3:], " ")
		AddTask(tasks, id, description)
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}

		id := os.Args[2]
		description := strings.Join(os.Args[3:], " ")
		UpdateTask(tasks, id, description)
	case "delete":
		if len(os.Args) < 2 {
			fmt.Println("Please provide a task id.")
			return
		}

		id := os.Args[2]
		DeleteTask(tasks, id)
	case "list":
		var filter string

		if len(os.Args) == 2 {
			ListTasks(tasks, "")
		} else {
			filter = os.Args[2]
			ListTasks(tasks, filter)
		}
	case "mark-in-progress", "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task id.")
			return
		}

		status := strings.Split(os.Args[1], "-")
		joinedStatus := strings.Join(status[1:], "-")
		id := os.Args[2]
		MarkTaskStatus(tasks, id, joinedStatus)
	default:
		println("Invalid operation. Use add, update, delete, or list.")
	}

}

func InitialReadTask() []Task {
	fileContent, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks file, creating the tasks.json file . . .")

		_, err := os.Create("tasks.json")
		if err != nil {
			fmt.Println("Error creating tasks file:", err.Error())
			return nil
		}

		fmt.Println("Success creating the tasks.json file")
		return nil
	}

	var tasks []Task
	err = json.Unmarshal(fileContent, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks JSON:", err.Error())
		return nil
	}

	return tasks
}

func AddTask(tasks []Task, id string, description string) {
	newTask := Task{
		Id:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	tasks = append(tasks, newTask)
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err.Error())
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err.Error())
		return
	}

	fmt.Println("Task added successfully!")
}

func UpdateTask(tasks []Task, id string, description string) {
	isFound := false
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			isFound = true
			break
		}
	}

	if !isFound {
		fmt.Println("Task with the given id not found.")
		return
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err.Error())
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err.Error())
		return
	}

	fmt.Println("Task updated successfully!")
}

func DeleteTask(tasks []Task, id string) {
	isFound := false
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Println("Task with the given id not found.")
		return
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err.Error())
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err.Error())
		return
	}

	fmt.Println("Task deleted successfully!")
}

func ListTasks(tasks []Task, filter string) {
	if len(tasks) == 0 || tasks == nil {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
				task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func MarkTaskStatus(tasks []Task, id string, status string) {
	isFound := false
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			isFound = true
			break
		}
	}

	if !isFound {
		fmt.Println("Task with the given id not found.")
		return
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err.Error())
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err.Error())
		return
	}

	fmt.Println("Task status updated successfully!")
}
