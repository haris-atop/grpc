package main

import (
	"context"
	"log"

	pb "github.com/haris-atop/Code/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked: %v", in)
	return &pb.GreetResponse{
		Result: "Hi " + in.FirstName,
	}, nil
}
