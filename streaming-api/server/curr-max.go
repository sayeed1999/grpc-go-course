package main

import (
	"io"
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
)

func (s *Server) GetCurrentMaximum(stream pb.PrimesService_GetCurrentMaximumServer) error {

	var max int32 = -99999

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("number received - ", req.Num)

		time.Sleep(time.Millisecond * 1000)
		if req.Num > max {
			max = req.Num
		}
		log.Println("max calculated - ", max)
		err = stream.Send(&pb.MaximumResponse{Num: max})
		if err != nil {
			return err
		}
	}
	return nil
}
