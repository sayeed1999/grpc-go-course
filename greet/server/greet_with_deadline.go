package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadlines(ctx context.Context, c *pb.GreetRequest) (*pb.GreetResponse, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Errorf(codes.Canceled, "The client canceled the request!")
		}
		time.Sleep(1 * time.Second)
	}

	message := "Welcome, " + c.FirstName + "!"
	log.Println("Sending:- ", message)
	return &pb.GreetResponse{
		Result: message,
	}, nil
}
