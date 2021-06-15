package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "v1"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	appHost := "localhost"
	appPort := 9000
	if port := os.Getenv("PORT"); port != "" {
		appPort = int(appPort)
	}

	if host := os.Getenv("HOST"); host != "" {
		appHost = host
	}

	log.Println("Starting grpc server")
	address := fmt.Sprintf("%s:%d", appHost, appPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBookServiceServer(grpcServer, &pb.BooksEndpoint{})

	log.Printf("Server running on %s\n", address)
	if grpcServer.Serve(lis) != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
