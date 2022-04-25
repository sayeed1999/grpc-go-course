package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
)

func getPrimesInRange(client pb.PrimesServiceClient) {
	time.Sleep(time.Second * 3)

	req := &pb.PrimesRequest{
		Num: 51,
	}

	stream, err := client.GetAllPrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while receiving data from stream: %v\n", err)
	}

	primes := []int64{}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		primes = append(primes, msg.Prime)
		log.Println("Primes in range 1-51: ", primes)
	}

	log.Println("printed all primes of ", req.Num)

	log.Print("Press enter to close the window.")
	fmt.Scanln()

	wg.Done()
}
