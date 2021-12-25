package main

import (
	"context"
	"fmt"
	"github.com/udodinho/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Calculator Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Created client: %f", c)
	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	request := &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 40,
	}
	response, err := c.Sum(context.Background(), request)
	if err != nil {
		fmt.Printf("Error while calling Sum RPC: %v", err)
	}
	fmt.Printf("Response from Sum: %v", response.Sum)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	request := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12087649,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling PimeDecomposition RPC: %v", err)
	}
	for {
		resStream, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(resStream.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client ComputeAverage Streaming RPC...")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream %v", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		fmt.Printf("Sending number %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}
	fmt.Printf("The average is: %v\n", res.GetAverage())
}
