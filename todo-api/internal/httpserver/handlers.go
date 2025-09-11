package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/example/todo-api/internal/auth"
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
// @Security BearerAuth
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
	id, err := s.Create(r.Context(), t)
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
// @Security BearerAuth
func handleListTodos(w http.ResponseWriter, r *http.Request, s repo.TodoRepository) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	todos, err := s.List(r.Context())
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
// @Security BearerAuth
func handleGetTodo(w http.ResponseWriter, r *http.Request, s repo.TodoRepository, id int) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	todo, err := s.Get(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get todo", http.StatusInternalServerError)
		fmt.Printf("Error getting todo: %v\n", err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// HandleUserRegistration godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User registration data"
// @Success 201 {object} map[string]string
// @Router /users/register [post]
func handleUserRegistration(w http.ResponseWriter, r *http.Request, s repo.TodoRepository) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var input model.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Printf("Error registering user: %v\n", err)
		return
	}

	userID, err := s.RegisterUser(r.Context(), &input)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		fmt.Printf("Error registering user: %v\n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": userID})
}

// HandleUserAuthentication godoc
// @Summary Authenticate a user
// @Description Authenticate a user with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserInput true "User login data"
// @Success 200 {object} map[string]string
// @Router /users/login [post]
func handleUserAuthentication(w http.ResponseWriter, r *http.Request, s repo.TodoRepository) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var input model.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Printf("Error authenticating user: %v\n", err)
		return
	}
	userID, err := s.AuthenticateUser(r.Context(), input.Username, input.Password)
	if err != nil {
		http.Error(w, "Failed to authenticate user", http.StatusUnauthorized)
		fmt.Printf("Error authenticating user: %v\n", err)
		return
	}
	strUsrID := fmt.Sprintf("%d", userID)
	token, err := auth.GenerateJWT(strUsrID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Printf("Error generating token: %v\n", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
