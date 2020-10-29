package main

import (
	"context"
	"log"
	"net/http"

	"github.com/kdelalic/go-microservice-structure/pkg/config"
	"github.com/kdelalic/go-microservice-structure/pkg/helpers"
	"github.com/kdelalic/go-microservice-structure/services/apigateway/server"
	"github.com/rs/cors"
)

func main() {
	// Create a new gateway server.
	gateway, err := server.CreateGateway(context.Background())
	if err != nil {
		log.Fatalf("Error in creating the gateways from other microservice protos : %v", err)
	}

	// Create a new ServeMux.
	mux := http.NewServeMux()

	// Registers the handler for the given pattern.
	mux.Handle("/", gateway)

	// TODO: Enable cors only for our frontend
	handler := cors.AllowAll().Handler(mux)

	log.Printf("Apigateway deployed on port %s\n", config.ApigatewayPort)

	go func() {
		err = http.ListenAndServe(":8080", handler)
		if err != nil {
			log.Fatalf("Error creating an HTTPS connection : %v", err)
		}
	}()

	// listens on the TCP network address and then calls
	// Serve with handler to handle requests on incoming HTTPS connections.
	err = http.ListenAndServeTLS(config.ApigatewayPort, helpers.SSLCertPath(config.ApigatewayName), helpers.SSLKeyPath(config.ApigatewayName), handler)
	if err != nil {
		log.Fatalf("Error creating an HTTPS connection : %v", err)
	}
}
