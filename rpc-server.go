// Package main implements a server for chartDataExchange service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yoeria/goftxlogistics/protos"
)

const (
	port = ":50051"
)

func (s *pb.ChartDataExchangeServer) serveWindowDataArray(ctx context.Context, in *pb.) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.WindowDataArray{}, nil
}

func (s *server) handleChartDataRequest(ctx context.Context, in *pb.ChartDataRequest) (*pb.ChartDataReply, error) {
}

func server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	grpc.
		log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
