package main

import (
	"fmt"
	"net"
	"todo-grpc/internal/server"
	"todo-grpc/internal/storage"
	todo_v1 "todo-grpc/proto/todo/v1"

	"google.golang.org/grpc"
)

func main() {
	grpcLis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer grpcLis.Close()

	repo := storage.NewInMemoryTodoRepo()
	grpcServer := grpc.NewServer()
	todo_v1.RegisterTodoServiceServer(grpcServer, server.NewTodoServer(repo))
	fmt.Print("gRPC server listening on :50051\n")
	if err := grpcServer.Serve(grpcLis); err != nil {
		panic(err)
	}

}
