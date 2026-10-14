package main

import (
	"context"
	"fmt"
	"log"
	"net"

	todo "github.com/sambacarlson/learning-grpc/todo_app/proto_gen"
	"google.golang.org/grpc"
)

type TodoServer struct {
	todo.UnimplementedTodoServiceServer
	todoDB []*todo.Todo
}

// Check that TodoServer implements the TodoServiceServer interface
var _ todo.TodoServiceServer = (*TodoServer)(nil)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	todoDB := []*todo.Todo{}
	todoDB = append(todoDB, &todo.Todo{Id: 0, Title: "Default todo", Category: todo.Category_CATEGORY_NORMAL})

	todo.RegisterTodoServiceServer(server, &TodoServer{todoDB: todoDB})

	fmt.Println("Server starting on 50051")
	if err := server.Serve(lis); err != nil {
		fmt.Println("Server failed to serve")
		log.Fatal(err)
	}
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.Todo, error) {
	s.todoDB = append(s.todoDB, &todo.Todo{
		Id:       int64(len(s.todoDB)),
		Title:    req.Title,
		Category: req.Category,
	})
	fmt.Printf("todo created: %v\n", req)
	return s.todoDB[len(s.todoDB)-1], nil
}

func (s *TodoServer) GetTodos(ctx context.Context, req *todo.GetTodosRequest) (*todo.GetTodosResponse, error) {
	fmt.Println("a request was made to get todos")
	return &todo.GetTodosResponse{Todos: s.todoDB}, nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, req *todo.UpdateTodoRequest) (*todo.Todo, error) {
	fmt.Printf("a request was made to update a todo item: %+v\n", req)
	return nil, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	fmt.Printf("a request was made to delete a todo item: %+v\n", req)
	return nil, nil
}
