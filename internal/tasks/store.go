package tasks

import (
	"errors"
	"strings"
	"sync"
)

var ErrEmptyTitle = errors.New("task title cannot be empty")

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

// ListPage is a draft API. The final behavior for invalid bounds is undecided.
func (s *Store) ListPage(offset, limit int) []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset >= len(s.tasks) {
		return nil
	}
	end := min(offset+limit, len(s.tasks))
	result := make([]Task, end-offset)
	copy(result, s.tasks[offset:end])
	return result
}
