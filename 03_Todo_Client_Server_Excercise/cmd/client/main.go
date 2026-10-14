package main

import (
	"context"
	"fmt"
	"log"

	todo "github.com/sambacarlson/learning-grpc/todo_app/proto_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := todo.NewTodoServiceClient(conn)

	ctx := context.Background()

	// create todo
	createTodoReq := &todo.CreateTodoRequest{
		Title:    "My cool todo",
		Category: todo.Category_CATEGORY_NORMAL,
	}
	todo, err := client.CreateTodo(ctx, createTodoReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todo)

	// get todos
	todos, err := client.GetTodos(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todos)

}
