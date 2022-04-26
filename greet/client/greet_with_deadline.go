package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Sayeed",
	}
	res, err := c.GreetWithDeadlines(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeed !")
				return
			} else {
				log.Fatalf("Unexcepted gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", res.Result)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
