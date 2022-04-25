// go build -o ./bin/streaming-api/server.exe ./streaming-api/server/.

package main

import (
	"log"
	"net"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PrimesServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening to %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterPrimesServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
