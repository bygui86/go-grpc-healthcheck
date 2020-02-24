package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	
	proto "github.com/bygui86/go-grpc-healthcheck/domain"
)

const (
	port = ":50051"
	
	healthServiceName = "grpc.health.v1.server"
)

// GrpcServer is used to implement helloworld.GreeterServer.
type GrpcServer struct {
	proto.UnimplementedGreeterServer
}

func main() {
	log.Println("Starting gRPC server")
	
	log.Println("Creating gRPC server")
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, &GrpcServer{})
	
	log.Println("Creating gRPC health-check server")
	healthServer := health.NewServer()
	
	healthServer.SetServingStatus(healthServiceName, 1)
	healthpb.RegisterHealthServer(server, healthServer)
	
	log.Println("Starting TCP listener")
	listener, listErr := net.Listen("tcp", port)
	if listErr != nil {
		log.Fatalf("failed to listen: %v", listErr)
	}
	log.Println("Starting gRPC servers")
	serverErr := server.Serve(listener)
	if serverErr != nil {
		log.Fatalf("failed to serve: %v", serverErr)
	}
}

// SayHello implements helloworld.GreeterServer
func (s *GrpcServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &proto.HelloReply{Message: "Hello " + in.GetName()}, nil
}
