package main

import (
	"log"
	"net"

	"math"

	pb "github.com/omustardo/demos/grpc-calculator/calc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.CalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.Numbers) (*pb.Result, error) {
	return &pb.Result{Value: in.X + in.Y}, nil
}
func (s *server) Sub(ctx context.Context, in *pb.Numbers) (*pb.Result, error) {
	return &pb.Result{Value: in.X - in.Y}, nil
}
func (s *server) Mul(ctx context.Context, in *pb.Numbers) (*pb.Result, error) {
	return &pb.Result{Value: in.X * in.Y}, nil
}
func (s *server) Div(ctx context.Context, in *pb.Numbers) (*pb.Result, error) {
	if in.Y == 0 {
		return &pb.Result{Value: math.NaN()}, nil
	}
	return &pb.Result{Value: in.X / in.Y}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
