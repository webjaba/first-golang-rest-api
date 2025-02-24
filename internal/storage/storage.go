package storage

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrTaskDoesNotExist = errors.New("task does not exist")
)

type Task struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done,omitempty"`
}

func CreateTask(jsonObj []byte) (Task, error) {
	t := Task{}

	err := json.Unmarshal(jsonObj, &t)

	if err != nil {
		return t, fmt.Errorf("unable to unmarshall json: %w", err)
	}

	return t, nil
}

func (t *Task) MarshallToJSON() ([]byte, error) {
	jsonObj, err := json.Marshal(*t)

	if err != nil {
		return []byte{}, fmt.Errorf("unable to marshall task: %w", err)
	}

	return jsonObj, nil
}
