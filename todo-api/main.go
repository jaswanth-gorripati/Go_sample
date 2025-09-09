package main

import (
	"fmt"
	"net/http"

	"github.com/example/todo-api/docs"
	"github.com/example/todo-api/internal/httpserver"
	"github.com/example/todo-api/internal/middleware"
	"github.com/example/todo-api/internal/repo"
	"github.com/example/todo-api/internal/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	fmt.Println("Todo API")

	s := &storage.Storage{
		Repository: repo.NewInMemoryTodoRepo(),
	}
	mux := httpserver.RegisterRoutes(s)
	docs.SwaggerInfo.BasePath = "/"
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	var handler http.Handler = mux
	handler = middleware.Recovery(handler)
	handler = middleware.Logging(handler)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}
	fmt.Println("Starting server on :8081")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
