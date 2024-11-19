package main

import (
	"fmt"
	"unit-test-adam/todo"
)

func main() {
	todoList := todo.TodoList{}

	todoList.AddTask("Belajar Golang")
	todoList.AddTask("Membuat aplikasi Todo")

	fmt.Println("Daftar Tugas:")
	for _, task := range todoList.GetTasks() {
		fmt.Printf("ID: %d, Task: %s, Completed: %t\n", task.ID, task.Task, task.Completed)
	}

	todoList.CompleteTask(1)
	fmt.Println("\nSetelah tugas pertama selesai:")
	for _, task := range todoList.GetTasks() {
		fmt.Printf("ID: %d, Task: %s, Completed: %t\n", task.ID, task.Task, task.Completed)
	}

	todoList.RemoveTask(2)
	fmt.Println("\nSetelah tugas kedua dihapus:")
	for _, task := range todoList.GetTasks() {
		fmt.Printf("ID: %d, Task: %s, Completed: %t\n", task.ID, task.Task, task.Completed)
	}
}
