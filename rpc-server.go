// Package main implements a server for chartDataExchange service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/yoeria/goftxlogistics/protos"
)

const (
	port = ":50051"
)

func (s *pb.ChartDataExchangeServer) serveWindowDataArray(ctx context.Context, in *pb.ChartDataExchangeClient) (*pb.WindowDataArray, error) {
	log.Printf("Received: %v")
	return &pb.WindowDataArray{}, nil
}

func handleChartDataRequest(ctx context.Context) {
}

func server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
