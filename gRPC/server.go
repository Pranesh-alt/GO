package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"user-service/models"
	"user-service/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type userServer struct {
	proto.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (s *userServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	user := models.User{Name: req.Name, Email: req.Email}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &proto.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	var user models.User
	if err := s.DB.First(&user, req.Id).Error; err != nil {
		return nil, err
	}
	return &proto.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect DB: ", err)
	}
	db.AutoMigrate(&models.User{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &userServer{DB: db})

	fmt.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
