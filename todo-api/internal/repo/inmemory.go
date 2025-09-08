package repo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/example/todo-api/internal/model"
)

type InMemoryTodoRepo struct {
	mu    sync.RWMutex
	todos map[int]*model.Todo
	seq   atomic.Uint32
}

// TODO: Discuss Later
func NewInMemoryTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		todos: make(map[int]*model.Todo),
		seq:   atomic.Uint32{},
	}
}

func (r *InMemoryTodoRepo) Create(todo *model.Todo) (string, error) {
	now := time.Now()
	rec := &model.Todo{
		ID:        int(r.seq.Load()),
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: now,
		UpdatedAt: now,
	}
	r.mu.Lock()
	r.todos[rec.ID] = rec
	r.mu.Unlock()
	r.seq.Add(1)
	return fmt.Sprintf("%d", rec.ID), nil
}
func (r *InMemoryTodoRepo) List() ([]*model.Todo, error) {
	return nil, nil
}
func (r *InMemoryTodoRepo) Get(id int) (*model.Todo, error) {
	return nil, nil
}
func (r *InMemoryTodoRepo) Update(todo *model.Todo) error {
	return nil
}
func (r *InMemoryTodoRepo) Delete(id int) error {
	return nil
}
func (r *InMemoryTodoRepo) SetDone(id int, done bool) (*model.Todo, error) {
	return nil, nil
}
