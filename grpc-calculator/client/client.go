package main

import (
	"fmt"
	"log"

	"github.com/omustardo/demos/grpc-calculator/calc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := calc.NewCalculatorClient(conn)

	// Contact the server and print out its response.
	inputs := &calc.Numbers{X: 1, Y: 2}
	sum, err := c.Add(context.Background(), inputs)
	if err != nil {
		log.Fatalln(err)
	}
	difference, err := c.Sub(context.Background(), inputs)
	if err != nil {
		log.Fatalln(err)
	}
	product, err := c.Mul(context.Background(), inputs)
	if err != nil {
		log.Fatalln(err)
	}
	quotient, err := c.Div(context.Background(), inputs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Results for input %v\n", inputs)
	fmt.Printf(" Sum: %v\n", sum.Value)
	fmt.Printf(" Difference: %v\n", difference.Value)
	fmt.Printf(" Product: %v\n", product.Value)
	fmt.Printf(" Quotient: %v\n", quotient.Value)
}
