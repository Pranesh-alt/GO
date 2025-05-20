package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/praneshragu/grpc/helloworldpb"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func (s *greeterServer) SayHelloManyTimes(req *pb.HelloRequest, stream pb.Greeter_SayHelloManyTimesServer) error {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello %s number %d", req.Name, i+1)
		stream.Send(&pb.HelloReply{Message: msg})
		time.Sleep(time.Second)
	}
	return nil
}

func (s *greeterServer) LongGreet(stream pb.Greeter_LongGreetServer) error {
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ", ")})
		}
		if err != nil {
			return err
		}
		names = append(names, req.Name)
	}
}

func (s *greeterServer) GreetEveryone(stream pb.Greeter_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		msg := "Hi " + req.Name
		if err := stream.Send(&pb.HelloReply{Message: msg}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
