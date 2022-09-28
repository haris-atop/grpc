package main

import (
	"fmt"
	"log"

	pb "github.com/haris-atop/Code/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	fmt.Println("client connected")
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	DoGreet(c)
}
