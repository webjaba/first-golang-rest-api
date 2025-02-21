package inmemory

import (
	"todo/internal/storage"
)

type Storage struct {
	Tasks map[int]storage.Task
}

func NewStorage() *Storage {
	return &Storage{Tasks: make(map[int]storage.Task)}
}

func (s *Storage) AddOrReplaceTask(jsonObj []byte) error {
	task, err := storage.CreateTask(jsonObj)

	if err != nil {
		return err
	}

	s.Tasks[task.ID] = task

	return nil
}

func (s *Storage) GetTask(id int) (*storage.Task, error) {
	task, exists := s.Tasks[id]

	if !exists {
		return nil, storage.ErrTaskDoesNotExist
	}

	return &task, nil
}

func (s *Storage) GetAllTasks() []storage.Task {
	result := []storage.Task{}

	for _, task := range s.Tasks {
		result = append(result, task)
	}

	return result
}

func (s *Storage) DeleteTask(id int) error {
	_, exists := s.Tasks[id]

	if !exists {
		return storage.ErrTaskDoesNotExist
	}

	delete(s.Tasks, id)

	return nil
}
