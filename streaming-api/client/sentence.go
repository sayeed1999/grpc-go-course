package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
)

func getWholeSentence(server pb.PrimesServiceClient) {
	time.Sleep(time.Second * 3)

	reqs := []*pb.WordRequest{
		{Word: "Hi"},
		{Word: "Hi"},
		{Word: "Hi"},
		{Word: "I"},
		{Word: "I"},
		{Word: "I"},
		{Word: "like"},
		{Word: "like"},
		{Word: "like"},
		{Word: "to"},
		{Word: "to"},
		{Word: "to"},
		{Word: "EAT"},
		{Word: "EAT"},
		{Word: "EAT"},
	}

	stream, err := server.GetWholeSentenceFromBrokenWords(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, req := range reqs {
		stream.Send(req)
		log.Println("sent - ", req.Word)
		time.Sleep(time.Second * 1)
	}

	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("The whole sentence is:- ", response.Sentence)

	log.Println("press enter to exit")
	fmt.Scanln()

	wg.Done()
}
