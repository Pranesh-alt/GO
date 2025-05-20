package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/praneshragu/grpc/helloworldpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// callSayHello(client)
	// callSayHelloManyTimes(client)
	callLongGreet(client)
	// callGreetEveryone(client)
}

func callSayHello(client pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Unary Greeting: %s", resp.Message)
}

func callSayHelloManyTimes(client pb.GreeterClient) {
	req := &pb.HelloRequest{Name: "Alice"}
	stream, err := client.SayHelloManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling SayHelloManyTimes: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		log.Printf("Server Stream: %s", resp.Message)
	}
}

func callLongGreet(client pb.GreeterClient) {
	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error calling LongGreet: %v", err)
	}

	names := []string{"Iron Man", "Friday", "Jarvis"}
	for _, name := range names {
		log.Printf("Sending: %s", name)
		stream.Send(&pb.HelloRequest{Name: name})
		time.Sleep(time.Millisecond * 300)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response from LongGreet: %v", err)
	}
	log.Printf("Client Stream Response: %s", resp.Message)
}

func callGreetEveryone(client pb.GreeterClient) {
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error calling GreetEveryone: %v", err)
	}

	waitc := make(chan struct{})

	// Receive responses
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving: %v", err)
			}
			log.Printf("Received: %s", resp.Message)
		}
		close(waitc)
	}()

	// Send requests
	names := []string{"Batman", "Robin", "Alfred"}
	for _, name := range names {
		log.Printf("Sending: %s", name)
		stream.Send(&pb.HelloRequest{Name: name})
		time.Sleep(time.Millisecond * 500)
	}
	stream.CloseSend()
	<-waitc
}
