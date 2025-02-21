package main

import (
	"fmt"
	"net/http"
	"todo/internal/config"
	"todo/internal/handlers"
	inmemory "todo/internal/storage/in_memory"
)

func main() {
	cfg := config.New()

	fmt.Println(cfg)

	_ = cfg

	storage := inmemory.NewStorage()

	http.HandleFunc("/", handlers.CreateTaskHandler(storage))

	http.ListenAndServe(":8080", nil)
}
