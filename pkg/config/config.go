package config

import (
	"github.com/kdelalic/go-microservice-structure/pkg/helpers"
)

// ApigatewayName is the name of apigateway
var ApigatewayName string = "apigateway"

// ApigatewayPort is the port that apigateway is deployed on
var ApigatewayPort string = helpers.GetEnv("APIGATEWAY_PORT", ":443")

// Service1ServiceName is the name of service1
var Service1ServiceName string = "service1"

// Service1ServicePort is the port that service1 is deployed on
var Service1ServicePort string = helpers.GetEnv("SERVICE1_PORT", ":10001")

// Service1ServiceEndpoint is the endpoint which service1 can be called from
var Service1ServiceEndpoint string = helpers.GetEnv("SERVICE1_ENDPOINT", "localhost"+Service1ServicePort)

//Service2ServiceName is the name of service2
var Service2ServiceName string = "service2"

// Service2ServicePort is the port that service2 is deployed on
var Service2ServicePort string = helpers.GetEnv("SERVICE2_PORT", ":10002")

// Service2ServiceEndpoint is the endpoint which service2 can be called from
var Service2ServiceEndpoint string = helpers.GetEnv("SERVICE2_ENDPOINT", "localhost"+Service2ServicePort)

// Environment is use to determine whether we are in production or development mode
var Environment string = helpers.GetEnv("GO_ENV", "development")
