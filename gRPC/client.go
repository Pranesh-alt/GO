package main

import (
	"context"
	"google.golang.org/grpc"
	"your/proto/package/path"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Duplicate declaration removed
	if err != nil {
		panic(err)
	}

	// Handle the response
	_ = resp
}
client := proto.NewUserServiceClient(conn)

resp, err := client.CreateUser(context.Background(), &proto.CreateUserRequest{
	Name:  "John",
	Email: "john@example.com",
})
