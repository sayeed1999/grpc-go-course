package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(server pb.PrimesServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := server.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server:%s\n", e.Message())
			log.Printf("Error code from server:%s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				time.Sleep(time.Second * 3)
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
	time.Sleep(time.Second * 3)
}
