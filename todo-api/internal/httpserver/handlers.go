package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/example/todo-api/internal/model"
	"github.com/example/todo-api/internal/repo"
)

func handleCreateTodo(w http.ResponseWriter, r *http.Request, s repo.TodoRepository) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var input model.TodoInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Printf("Error creating todo: %v\n", err)
		return
	}
	t := &model.Todo{Title: input.Title, Done: false}
	id, err := s.Create(t)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		fmt.Printf("Error creating todo: %v\n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
