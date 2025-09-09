package httpserver

import (
	"net/http"
	"strconv"

	"github.com/example/todo-api/internal/storage"
)

func RegisterRoutes(s *storage.Storage) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreateTodo(w, r, s.Repository)
		case http.MethodGet:
			handleListTodos(w, r, s.Repository)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/todos/"):]
		if idStr == "" {
			http.Error(w, "Missing todo ID", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid todo ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			handleGetTodo(w, r, s.Repository, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
