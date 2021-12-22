package main

import (
	"context"
	"fmt"
	"github.com/udodinho/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Received Sum RPC: %v", req)
	num1 := req.Num1
	num2 := req.Num2

	sum := num1 + num2
	res := &calculatorpb.SumResponse{
		Sum: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased %v\n", divisor)
		}
	}
	return nil
}

func main() {
	fmt.Println("Calculator Server")

	listerner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listerner); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
