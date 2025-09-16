package main

import (
	"context"
	"fmt"
	"io"
	"time"
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

	go func() {
		stream, err := c.WatchTodos(ctx)
		if err != nil {
			panic(err)
		}
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}
			fmt.Println()
			fmt.Printf("Received WatchTodos event: %+v\n", resp)
			fmt.Println()
		}
	}()

	_, err = c.CreateTodo(ctx, &todo_v1.CreateTodoRequest{Title: "My First Todo"})
	if err != nil {
		panic(err)
	}
	bulkStream, err := c.BulkCreateTodos(ctx)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		err := bulkStream.Send(&todo_v1.CreateTodoRequest{Title: fmt.Sprintf("Bulk Todo %d", i+1)})
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
	_, err = bulkStream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Bulk Created Todo Ids: %v\n", respBulkTodo.Ids)

	stream, err := c.ListTodos(ctx, &todo_v1.ListTodosRequest{Completed: false})
	if err != nil {
		panic(err)
	}
	for {
		_, err := stream.Recv()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		// fmt.Printf("Received Todo: %+v\n", todo)
	}
	time.Sleep(1 * time.Minute)
}
