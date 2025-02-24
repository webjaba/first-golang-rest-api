package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todo/internal/storage"
)

type Storage interface {
	AddOrReplaceTask(jsonObj []byte) error
	GetTask(id int) (*storage.Task, error)
	GetAllTasks() []storage.Task
	DeleteTask(id int) error
}

func CreateTaskHandler(s Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			fmt.Println("POST")

			body, err := io.ReadAll(r.Body)

			if err != nil {
				fmt.Println("Error of reading body:", err)
				w.WriteHeader(http.StatusBadRequest)
			}

			err = s.AddOrReplaceTask(body)

			if err != nil {
				fmt.Println("Error of creation task:", err)
				w.WriteHeader(http.StatusBadRequest)
			}
		} else if r.Method == "GET" {
			fmt.Println("GET")

			jsonData, err := json.Marshal(s.GetAllTasks())

			if err != nil {
				fmt.Println("Error of marshaling tasks", err)
				w.WriteHeader(http.StatusBadGateway)
			}

			w.Write(jsonData)

		} else if r.Method == "DELETE" {
			fmt.Println("delete")
		}
	}
}
