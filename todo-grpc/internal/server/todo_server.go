package server

import (
	"context"
	"fmt"
	"io"
	"sync"
	"todo-grpc/internal/storage"
	todo_v1 "todo-grpc/proto/todo/v1"
)

type SrvTd struct {
	todo_v1.UnimplementedTodoServiceServer
	repo *storage.InMemoryTodoRepo

	subsMu sync.Mutex
	subs   map[int]chan *todo_v1.WatchTodosResponse
	nextID int
}

func NewTodoServer(repo *storage.InMemoryTodoRepo) *SrvTd {
	return &SrvTd{
		repo: repo,
		subs: make(map[int]chan *todo_v1.WatchTodosResponse),
	}
}

func (s *SrvTd) WatchOPSEvents(ev *todo_v1.WatchTodosResponse) {
	s.subsMu.Lock()
	defer s.subsMu.Unlock()

	for _, ch := range s.subs {
		select {
		case ch <- ev:
		default:
		}
	}

}

func (s *SrvTd) CreateTodo(ctx context.Context, req *todo_v1.CreateTodoRequest) (*todo_v1.CreateTodoResponse, error) {
	fmt.Printf("Received CreateTodo request: %+v\n", req)
	todo, err := s.repo.Create(ctx, req.Title)
	if err != nil {
		return nil, err
	}
	s.WatchOPSEvents(&todo_v1.WatchTodosResponse{Todo: todo, Operation: todo_v1.WatchTodosOperation_CREATED})
	return &todo_v1.CreateTodoResponse{
		Todo: todo,
	}, nil
}

func (s *SrvTd) ListTodos(req *todo_v1.ListTodosRequest, stream todo_v1.TodoService_ListTodosServer) error {
	todos, err := s.repo.List(stream.Context(), req.Completed)
	if err != nil {
		return err
	}
	for _, todo := range todos {
		if err := stream.Send(todo); err != nil {
			return err
		}
		//time.Sleep(2 * time.Second)
		s.WatchOPSEvents(&todo_v1.WatchTodosResponse{Todo: todo, Operation: todo_v1.WatchTodosOperation_FETCHED})
	}
	return nil
}
func (s *SrvTd) BulkCreateTodos(stream todo_v1.TodoService_BulkCreateTodosServer) error {
	ids := []string{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&todo_v1.BulkCreateTodosResponse{
				Ids: ids,
			})
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received BulkCreateTodos request: %v\n", req.Title)
		todo, err := s.repo.Create(stream.Context(), req.Title)
		if err != nil {
			return err
		}
		ids = append(ids, todo.Id)
		s.WatchOPSEvents(&todo_v1.WatchTodosResponse{Todo: todo, Operation: todo_v1.WatchTodosOperation_CREATED})
	}
}

func (s *SrvTd) WatchTodos(stream todo_v1.TodoService_WatchTodosServer) error {
	s.subsMu.Lock()
	id := s.nextID
	s.nextID++
	ch := make(chan *todo_v1.WatchTodosResponse, 10)
	s.subs[id] = ch
	s.subsMu.Unlock()
	go func() {
		for {
			if _, err := stream.Recv(); err != nil {
				return
			}
		}
	}()

	for v := range ch {
		if err := stream.Send(v); err != nil {
			return err
		}
	}
	return nil
}
