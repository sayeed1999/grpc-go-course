package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("sqrt was invoked with: %v\n", in)

	number := in.Number
	if number < 0 {
		return nil,
			status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("Received a negative number: %d\n", number),
			)
	}

	result := math.Sqrt(float64(number))
	log.Println("sending result - ", result)
	time.Sleep(time.Second * 3)
	return &pb.SqrtResponse{
		Result: result,
	}, nil
}
