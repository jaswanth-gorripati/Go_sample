package storage

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	todo_v1 "todo-grpc/proto/todo/v1"
)

type InMemoryTodoRepo struct {
	mu    sync.RWMutex
	todos map[int]*todo_v1.Todo
	seq   atomic.Uint32
}

// TODO: Discuss Later
func NewInMemoryTodoRepo() *InMemoryTodoRepo {
	return &InMemoryTodoRepo{
		todos: make(map[int]*todo_v1.Todo),
		seq:   atomic.Uint32{},
	}
}

func (r *InMemoryTodoRepo) Create(ctx context.Context, title string) (*todo_v1.Todo, error) {
	now := time.Now()
	rec := &todo_v1.Todo{
		Id:        fmt.Sprintf("%d", r.seq.Load()),
		Title:     title,
		Completed: false,
		CreatedAt: int32(now.Unix()),
		UpdatedAt: int32(now.Unix()),
	}
	r.mu.Lock()
	intID, err := strconv.Atoi(rec.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID")
	}
	r.todos[intID] = rec
	r.mu.Unlock()
	r.seq.Add(1)
	return rec, nil
}
func (r *InMemoryTodoRepo) List(ctx context.Context, onlyTrue bool) ([]*todo_v1.Todo, error) {
	out := make([]*todo_v1.Todo, 0, len(r.todos))
	r.mu.RLock()
	for _, v := range r.todos {
		// if onlyTrue && !v.Completed {
		// 	continue
		// }
		out = append(out, v)
	}
	r.mu.RUnlock()
	return out, nil
}
func (r *InMemoryTodoRepo) Get(ctx context.Context, id int) (*todo_v1.Todo, error) {
	r.mu.RLock()
	todo, exists := r.todos[id]
	r.mu.RUnlock()
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}
	return todo, nil
}
func (r *InMemoryTodoRepo) Update(ctx context.Context, todo *todo_v1.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	intID, err := strconv.Atoi(todo.Id)
	if err != nil {
		return fmt.Errorf("invalid ID")
	}
	existing, exists := r.todos[intID]
	if !exists {
		return fmt.Errorf("todo not found")
	}
	existing.Title = todo.Title
	existing.Completed = todo.Completed
	existing.UpdatedAt = int32(time.Now().Unix())
	return nil
}
func (r *InMemoryTodoRepo) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.todos, id)
	return nil
}
func (r *InMemoryTodoRepo) SetDone(ctx context.Context, id int, done bool) (*todo_v1.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, exists := r.todos[id]
	if !exists {
		return nil, fmt.Errorf("todo not found")
	}
	existing.Completed = done
	existing.UpdatedAt = int32(time.Now().Unix())
	return existing, nil
}
