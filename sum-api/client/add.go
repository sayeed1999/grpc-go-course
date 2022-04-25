package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sayeed1999/grpc-go-course/sum-api/proto"
)

var x int32 = 0
var y int32 = 2

func doAdd(client pb.SumServiceClient) {
	log.Print("doAdd on the client invoked..")
	res, err := client.Add(context.Background(), &pb.SumRequest{
		Param1: x,
		Param2: y,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum returned:- %d\n", res.Result)
	x += 2
	y += 2
}
