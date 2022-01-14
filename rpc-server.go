// Package main implements a server for chartDataExchange service.
package main

import (
	"log"
	"net"

	pb "github.com/yoeria/goftxlogistics/protos"
	"google.golang.org/grpc"
)

var srv *pb.ChartDataExchangeServer
var s = grpc.NewServer()
var serviceRegistrar grpc.ServiceRegistrar

const (
	port = ":50051"
)

/* func serveWindowDataArray(ctx context.Context, in *pb.ChartDataExchangeClient) (*pb.WindowDataArray, error) {
	log.Printf("Received: %v")
	return &pb.WindowDataArray{}, nil
} */

func server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	pb.RegisterChartDataExchangeServer(serviceRegistrar, *srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
