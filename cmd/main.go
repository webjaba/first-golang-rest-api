package main

import (
	"net/http"
	"todo/internal/handlers"
	inmemory "todo/internal/storage/in_memory"
)

func main() {
	storage := inmemory.NewStorage()

	http.HandleFunc("/createtask", handlers.CreateTaskHandler(storage))

	http.ListenAndServe(":8080", nil)
}
