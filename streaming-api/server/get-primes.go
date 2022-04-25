package main

import (
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
)

func (s *Server) GetAllPrimes(req *pb.PrimesRequest, stream pb.PrimesService_GetAllPrimesServer) error {

	var i int64
	var err error

	log.Println("number received - ", req.Num)

	if req.Num >= 2 {
		time.Sleep(time.Millisecond * 800)
		log.Println("sending prime - ", 2)
		err = stream.Send(&pb.PrimesResponse{Prime: 2})
		if err != nil {
			return err
		}
	}

	for i = 3; i <= req.Num; i += 2 {
		if isPrime(i) == true {
			log.Println("sending prime - ", i)
			time.Sleep(time.Millisecond * 800)
			if err = stream.Send(&pb.PrimesResponse{
				Prime: i,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

// not efficient algorithm
func isPrime(n int64) bool {
	if n == 2 {
		return true
	}
	var i int64
	for i = 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
