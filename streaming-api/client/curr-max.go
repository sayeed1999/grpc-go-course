package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
)

func findCurrentMaximum(server pb.PrimesServiceClient) {
	defer wg.Done()
	res := []int32{1, 4, 2, 5, 3, 6, 9, 8}
	var max int32 = -99999
	wg2 := &sync.WaitGroup{}
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)

	stream, err := server.GetCurrentMaximum(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	// client sending
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		time.Sleep(time.Second * 2)

		for _, req := range res {
			time.Sleep(time.Millisecond * 1500)
			stream.Send(&pb.MaximumRequest{Num: req})
			log.Printf("sent - %v\n", req)
			ch <- 1 // send
			<-ch2   // receive from ch2
		}
		stream.CloseSend()
	}()

	// client receiving
	// wg2.Add(1) since this is an infinite loop, nobody should wait for this one to finish. let the main program terminate.
	go func() {
		// defer wg2.Done()

		for {
			<-ch // receive
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			max = req.Num
			log.Println("max till now - ", max)
			ch2 <- 2 // send from ch2
		}
	}()

	wg2.Wait()
}
