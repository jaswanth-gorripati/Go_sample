package httpserver

import (
	"net/http"

	"github.com/example/todo-api/internal/storage"
)

func RegisterRoutes(s *storage.Storage) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreateTodo(w, r, s.Repository)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
