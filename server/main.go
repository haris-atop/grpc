// this is grpc gateway

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/haris-atop/Code/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.TesServiceServer
}

func (s *Server) Tes1(ctx context.Context, req *pb.TesRequest) (*pb.TesRequest, error) {

	return req, nil
}
func main() {
	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterTesServiceHandlerServer(context.Background(), mux, &Server{})
		fmt.Println("Http Server Starts")
		log.Fatalln(http.ListenAndServe("localhost:8080", mux))
	}()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("err euy: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTesServiceServer(s, &Server{})
	fmt.Println("grpc Server Starts")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("grpc error euy: %v", err)
	}
}

// this is pure grpc server
// package main

// import (
// 	"log"
// 	"net"

// 	pb "github.com/haris-atop/Code/proto"
// 	"google.golang.org/grpc"
// )

// var addr string = "0.0.0.0:50051"

// type Server struct {
// 	pb.GreetServiceServer
// }
// func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
// 	log.Printf("Greet was invoked: %v", in)
// 	return &pb.GreetResponse{
// 		Result: "Hi " + in.FirstName,
// 	}, nil
// }
// func main() {
// 	lis, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		log.Fatalf("err euy: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreetServiceServer(s, &Server{})
// 	err = s.Serve(lis)
// 	if err != nil {
// 		log.Fatalf("grpc error euy: %v", err)
// 	}
// }
