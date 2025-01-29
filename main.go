package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"VIM", "Git", "Go", "Docker"}

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/tasks", getTasks)
	http.ListenAndServe(":8080", nil)
}

func getTasks(response http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintf(response, "%s\n", task)
	}
}

// func getTask() {
// 	fmt.Printf("Get single")
// }
