package main

import (
	"context"
	"log"
	"net"

	"github.com/sambacarlson/learning-grpc/proto_gen/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServer struct {
	users.UnimplementedUserServiceServer //All implementation of UserServiceServer(The API for UserService service) must embed UnimplementedUserServiceServer for forward compatibility.
}

func main() {
	// make client request to server on localhost:50051
	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //With TransTsportCredentials(insecure.NewCredentials()) creates a connection without TLS, useful for dev/testing or when you have a trusted network.
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := users.NewUserServiceClient(conn)

	resp, err := client.GetUser(ctx, &users.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("error occurred getting users; %v", err)
	}
	log.Printf("User: %v\n", resp)

	// Start a grpc server on localhost:50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	users.RegisterUserServiceServer(s, &UserServer{})
	log.Println("Server listening on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (s *UserServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.UserProfile, error) {
	user := &users.User{
		Id:    req.Id,
		Name:  "Alice",
		Roles: []users.RoleType{users.RoleType_ROLE_ADMIN},
	}
	return &users.UserProfile{User: user}, nil
}
