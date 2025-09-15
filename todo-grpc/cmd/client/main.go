package main

import (
	"context"
	"fmt"
	todo_v1 "todo-grpc/proto/todo/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// This is a placeholder main function.
	con, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer con.Close()
	c := todo_v1.NewTodoServiceClient(con)
	ctx := context.Background()
	todo, err := c.CreateTodo(ctx, &todo_v1.CreateTodoRequest{Title: "My First Todo"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Todo: %+v\n", todo.Todo)
}
