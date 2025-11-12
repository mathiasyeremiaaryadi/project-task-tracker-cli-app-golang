package main

import "os"

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func main() {
	operation := os.Args[1]
	switch operation {
	case "add":
		AddTask()
	case "update":
		UpdateTask()
	case "delete":
		DeleteTask()
	case "list":
		ListTasks()
	default:
		println("Invalid operation. Use add, update, delete, or list.")
	}

}

func InitialReadTask() {

}

func AddTask() {

}

func UpdateTask() {

}

func DeleteTask() {

}

func ListTasks() {

}
