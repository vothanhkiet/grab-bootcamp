package transport

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// RunGrpcServer runs gRPC service
func RunGrpcServer(ctx context.Context, service FeedbackServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	RegisterFeedbackServiceServer(server, service)

	// start gRPC server
	log.Println("Starting gRPC server...")
	return server.Serve(listen)
}
