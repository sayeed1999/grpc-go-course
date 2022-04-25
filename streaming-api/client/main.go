// To build client project, run:-
// go build -o ./bin/streaming-api/client.exe ./streaming-api/client/.

package main

import (
	"log"
	"sync"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"
var wg *sync.WaitGroup

func main() {
	wg = &sync.WaitGroup{} // initializes a wait group

	// grpc connection initialization...
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	// all connections closing at the end of program
	defer conn.Close()
	// initializing grpc client
	client := pb.NewPrimesServiceClient(conn)
	// waiting for two go routines to finish
	wg.Add(3)
	// first goroutine fired that is getting all primes in range
	go getPrimesInRange(client)
	// second goroutine fired that is making a sentence from broken words
	go getWholeSentence(client)
	// both goroutines have finished their task!
	go findCurrentMaximum(client)
	wg.Wait()
}
