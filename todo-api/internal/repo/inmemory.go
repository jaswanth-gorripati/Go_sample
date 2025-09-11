package repo

import (
	"context"
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

func (r *InMemoryTodoRepo) Create(ctx context.Context, todo *model.Todo) (string, error) {
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
func (r *InMemoryTodoRepo) List(ctx context.Context) ([]*model.Todo, error) {
	out := make([]*model.Todo, 0, len(r.todos))
	r.mu.RLock()
	for _, v := range r.todos {
		out = append(out, v)
	}
	r.mu.RUnlock()
	return out, nil
}
func (r *InMemoryTodoRepo) Get(ctx context.Context, id int) (*model.Todo, error) {
	r.mu.RLock()
	todo, exists := r.todos[id]
	r.mu.RUnlock()
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}
	return todo, nil
}
func (r *InMemoryTodoRepo) Update(ctx context.Context, todo *model.Todo) error {
	return nil
}
func (r *InMemoryTodoRepo) Delete(ctx context.Context, id int) error {
	return nil
}
func (r *InMemoryTodoRepo) SetDone(ctx context.Context, id int, done bool) (*model.Todo, error) {
	return nil, nil
}
