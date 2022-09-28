package main

import (
	"context"
	"log"

	pb "github.com/haris-atop/Code/proto"
)

func DoGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked from client")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "bambang",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting %v ", res.Result)
}
