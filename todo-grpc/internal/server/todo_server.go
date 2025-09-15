package server

import (
	"context"
	"fmt"
	"todo-grpc/internal/storage"
	todo_v1 "todo-grpc/proto/todo/v1"
)

type SrvTd struct {
	todo_v1.UnimplementedTodoServiceServer
	repo *storage.InMemoryTodoRepo
}

func NewTodoServer(repo *storage.InMemoryTodoRepo) *SrvTd {
	return &SrvTd{
		repo: repo,
	}
}

func (s *SrvTd) CreateTodo(ctx context.Context, req *todo_v1.CreateTodoRequest) (*todo_v1.CreateTodoResponse, error) {
	fmt.Printf("Received CreateTodo request: %+v\n", req)
	todo, err := s.repo.Create(ctx, req.Title)
	if err != nil {
		return nil, err
	}
	return &todo_v1.CreateTodoResponse{
		Todo: todo,
	}, nil
}
