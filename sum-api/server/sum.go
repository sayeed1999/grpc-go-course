package main

import (
	"context"
	"log"

	pb "github.com/sayeed1999/grpc-go-course/sum-api/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Print("Add method on server invoked!")
	return &pb.SumResponse{
		Result: req.Param1 + req.Param2,
	}, nil
}
