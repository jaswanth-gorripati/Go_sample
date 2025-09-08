package storage

import "github.com/example/todo-api/internal/repo"

type Storage struct {
	Repository repo.TodoRepository
	// DB *sql.DB
}
