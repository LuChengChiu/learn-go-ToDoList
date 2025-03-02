package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("### Golang main init")
}

var taskItem1 = "The main goal of this project is to learn Golang."
var taskItem2 = "The secondary goal is to have one more side project, even it's for learning purpose."
var taskItem3 = "Enjoy the process of developing, even it's a very simple app."
var taskItems = []string{taskItem1, taskItem2, taskItem3}

func main() {
	fmt.Println("###### Welcome to the ToDoList App for learning Golang !")

	fmt.Println()
	printTasks(taskItems)
	taskItems = addTask(taskItems, "Live a wonder life!")
	fmt.Println("New Todos after adding one")
	printTasks(taskItems)

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/tasks", renderTasks)
	http.ListenAndServe(":8080", nil)
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	greeting := "Hello! This is an to do list app for learning Golang."
	fmt.Fprintf(w, greeting)
}

func renderTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A List of my Todos: ")
	for index, task := range taskItems {
		fmt.Fprintf(w, "%d. %s \n", index+1, task)
	}
}

func printTasks(tasks []string) {
	fmt.Println("A List of my Todos: ")
	for index, task := range tasks {
		fmt.Printf("%d. %s \n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updatedTaskItems := append(taskItems, newTask)
	return updatedTaskItems
}
