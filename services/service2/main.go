package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/kdelalic/go-microservice-structure/pkg/config"
	"github.com/kdelalic/go-microservice-structure/pkg/helpers"
	pb "github.com/kdelalic/go-microservice-structure/pkg/proto/service2"
	"github.com/kdelalic/go-microservice-structure/services/service2/controllers"
)

func main() {
	lis, err := net.Listen("tcp", config.Service2ServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create server certificates.
	creds, err := credentials.NewServerTLSFromFile(helpers.SSLCertPath(config.Service2ServiceName), helpers.SSLKeyPath(config.Service2ServiceName))
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	// Create a new server using the created credentials.
	gRPCServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterService2Server(gRPCServer, &controllers.Service2Server{})

	log.Printf("Service2 deployed on: %s\n", config.Service2ServicePort)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
