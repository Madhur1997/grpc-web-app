package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-web-app/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) Add(context context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Println("Got a new Add request")
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2
	result := &calculatorpb.AddResponse{Result: sum}
	return result, nil
}

func main() {
	fmt.Println("Starting Calculator server")
	lis, err := net.Listen("tcp", "0.0.0.0:50551")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
