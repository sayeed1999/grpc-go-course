package main

import (
	"io"
	"log"

	pb "github.com/sayeed1999/grpc-go-course/streaming-api/proto"
)

func (s *Server) GetWholeSentenceFromBrokenWords(stream pb.PrimesService_GetWholeSentenceFromBrokenWordsServer) error {

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res += "!"
			return stream.SendAndClose(&pb.SentenceResponse{Sentence: res})
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println("word received - ", req.Word)
		res += req.Word + " "
	}

}
