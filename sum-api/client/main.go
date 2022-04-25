// To build client project, run:-
// go build -o ./bin/sum-api/client.exe ./sum-api/client/.

package main

import (
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/sum-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSumServiceClient(conn)

	for {
		doAdd(client)
		time.Sleep(time.Second * 2)
	}
}
