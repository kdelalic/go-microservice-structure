package server

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/kdelalic/go-microservice-structure/pkg/config"
	"github.com/kdelalic/go-microservice-structure/pkg/helpers"
	pb_service1 "github.com/kdelalic/go-microservice-structure/pkg/proto/service1"
	pb_service2 "github.com/kdelalic/go-microservice-structure/pkg/proto/service2"
)

// CreateGateway creates the gateway to access other microservices.
func CreateGateway(ctx context.Context, muxOptions ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(muxOptions...)

	var err error

	// Create client credentials with certificate if we are not on development
	var service1Credentials credentials.TransportCredentials
	if config.Environment == "development" {
		service1Credentials = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	} else {
		service1Credentials, err = credentials.NewClientTLSFromFile(helpers.SSLCertPath(config.Service1ServiceName), "")
		if err != nil {
			log.Fatalf("Error in creating server credentials. %v", err)
			return nil, err
		}
	}

	// Create service1 service dial options using the created credentials.
	service1DialOptions := []grpc.DialOption{grpc.WithTransportCredentials(service1Credentials)}

	// Register the service1 service handler from endpoint using the created dial options.
	err = pb_service1.RegisterService1HandlerFromEndpoint(ctx, mux, config.Service1ServiceEndpoint, service1DialOptions)
	if err != nil {
		log.Fatalf("Error in registering end point. %v", err)
		return nil, err
	}

	// Create client credentials with certificate if we are not on development
	var service2Credentials credentials.TransportCredentials
	if config.Environment == "development" {
		service2Credentials = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	} else {
		service2Credentials, err = credentials.NewClientTLSFromFile(helpers.SSLCertPath(config.Service2ServiceName), "")
		if err != nil {
			log.Fatalf("Error in creating server credentials. %v", err)
			return nil, err
		}
	}

	// Create service2 service dial options using the created credentials.
	service2DialOptions := []grpc.DialOption{grpc.WithTransportCredentials(service2Credentials)}

	// Register the service2 service handler from endpoint using the created dial options.
	err = pb_service2.RegisterService2HandlerFromEndpoint(ctx, mux, config.Service2ServiceEndpoint, service2DialOptions)
	if err != nil {
		log.Fatalf("Error in registering end point. %v", err)
		return nil, err
	}

	return mux, nil
}
