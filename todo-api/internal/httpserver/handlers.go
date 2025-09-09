package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/example/todo-api/internal/model"
	"github.com/example/todo-api/internal/repo"
)

// HandleCreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body model.TodoInput true "Todo item"
// @Success 201 {object} model.Todo
// @Router /todos [post]
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

// HandleListTodos godoc
// @Summary List all todos
// @Description Get a list of all todo items
// @Tags todos
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos [get]
func handleListTodos(w http.ResponseWriter, r *http.Request, s repo.TodoRepository) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	todos, err := s.List()
	if err != nil {
		http.Error(w, "Failed to list todos", http.StatusInternalServerError)
		fmt.Printf("Error listing todos: %v\n", err)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// HandleGetTodo godoc
// @Summary Get a todo by ID
// @Description Get a specific todo item by its ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} model.Todo
// @Router /todos/{id} [get]
func handleGetTodo(w http.ResponseWriter, r *http.Request, s repo.TodoRepository, id int) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	todo, err := s.Get(id)
	if err != nil {
		http.Error(w, "Failed to get todo", http.StatusInternalServerError)
		fmt.Printf("Error getting todo: %v\n", err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}
