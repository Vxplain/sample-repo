package tasks

import (
	"errors"
	"strings"
	"sync"
)

var ErrEmptyTitle = errors.New("task title cannot be empty")
var ErrTaskNotFound = errors.New("task not found")

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Store struct {
	mu     sync.RWMutex
	nextID int
	tasks  []Task
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (s *Store) Add(title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	task := Task{ID: s.nextID, Title: title}
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s *Store) List() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Task, len(s.tasks))
	copy(result, s.tasks)
	return result
}

func (s *Store) Complete(id int) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.tasks {
		if s.tasks[i].ID == id {
			updated := s.tasks[i]
			updated.Done = true
			s.tasks[i] = updated
			return updated, nil
		}
	}
	return Task{}, ErrTaskNotFound
}
