package main

import (
	"fmt"
	"net/http"

	"github.com/example/todo-api/internal/httpserver"
	"github.com/example/todo-api/internal/repo"
	"github.com/example/todo-api/internal/storage"
)

func main() {
	fmt.Println("Todo API")

	s := &storage.Storage{
		Repository: repo.NewInMemoryTodoRepo(),
	}
	mux := httpserver.RegisterRoutes(s)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	fmt.Println("Starting server on :8081")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
