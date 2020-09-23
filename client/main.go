package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	proto "github.com/bygui86/go-grpc-healthcheck/domain"
)

const (
	address     = "localhost:50051"
	defaultName = "world"

	healthServiceName = "grpc.health.v1.server"
)

func main() {
	log.Println("Starting gRPC client")

	log.Println("Creating gRPC connection")
	// Set up a connection to the server.
	grpcConnection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcConnection.Close()

	log.Println("Creating gRPC health-check client")
	healthClient := healthpb.NewHealthClient(grpcConnection)

	// log.Println("Creating gRPC client")
	// grpcClient := proto.NewGreeterClient(grpcConnection)

	log.Println("Creating gRPC client")
	// counter := 0
	for {
		// if counter%3 == 0 {
		go check(healthClient)
		time.Sleep(1 * time.Second)
		// }

		// go greet(grpcClient)
		// time.Sleep(3 * time.Second)
		// counter++
	}
}

func check(healthClient healthpb.HealthClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	request := &healthpb.HealthCheckRequest{
		Service: healthServiceName,
	}

	response, err := healthClient.Check(ctx, request)
	if err != nil {
		log.Fatalf("could not check: %s", err.Error())
	}
	log.Printf("Status: %s", response.Status.String())
}

func greet(grpcClient proto.GreeterClient) {
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Contact the server and print out its response.
	response, err := grpcClient.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %s", err.Error())
	}
	log.Printf("Greeting: %s", response.GetMessage())
}
