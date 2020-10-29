package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/kdelalic/go-microservice-structure/pkg/config"
	"github.com/kdelalic/go-microservice-structure/pkg/helpers"
	pb "github.com/kdelalic/go-microservice-structure/pkg/proto/service1"
	"github.com/kdelalic/go-microservice-structure/services/service1/controllers"
)

func main() {
	lis, err := net.Listen("tcp", config.Service1ServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create server certificates.
	creds, err := credentials.NewServerTLSFromFile(helpers.SSLCertPath(config.Service1ServiceName), helpers.SSLKeyPath(config.Service1ServiceName))
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	// Create a new server using the created credentials.
	gRPCServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterService1Server(gRPCServer, &controllers.Service1Server{})

	log.Printf("Service1 deployed on: %s\n", config.Service1ServicePort)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
