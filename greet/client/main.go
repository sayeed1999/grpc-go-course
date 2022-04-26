// go build -o ./bin/greet/greetwithdeadline.exe ./greet/client/.

package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/sayeed1999/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create connection: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	log.Println("greet before deadline...")
	doGreetWithDeadline(c, time.Second*5)
	log.Println("greet after deadline...")
	doGreetWithDeadline(c, time.Second*2)
	log.Println("greet on deadline, server will return but client will throw deadline...")
	doGreetWithDeadline(c, time.Second*3)
	fmt.Scanln()
}
