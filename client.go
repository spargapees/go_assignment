package main

import (
	pb "assignment/user" // Import the generated protocol buffer package
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	addUserResponse, err := client.AddUser(context.Background(), &pb.User{
		Id:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	})
	if err != nil {
		log.Fatalf("AddUser failed: %v", err)
	}
	fmt.Printf("AddUser response: %v\n", addUserResponse)

	getUserResponse, err := client.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	fmt.Printf("GetUser response: %v\n", getUserResponse)

	listUsersResponse, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("ListUsers failed: %v", err)
	}
	fmt.Printf("ListUsers response: %v\n", listUsersResponse)
}
