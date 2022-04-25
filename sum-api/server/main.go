// go build -o ./bin/sum-api/server.exe ./sum-api/server/.

package main

import (
	"log"
	"net"

	pb "github.com/sayeed1999/grpc-go-course/sum-api/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening to %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
