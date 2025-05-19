package main

type server struct {
	pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := User{Name: req.Name, Email: req.Email}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func main() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&User{})

	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{DB: db})

	fmt.Println("gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
