package controllers

// Load required packages
import (
	"context"

	model "github.com/kdelalic/go-microservice-structure/pkg/proto/service1"
	"github.com/kdelalic/go-microservice-structure/services/service1/services"
)

// Service1Server is the service1 server
type Service1Server struct {
}

// Func1 does something based on input
func (server *Service1Server) Func1(ctx context.Context, in *model.Func1Request) (*model.Func1Response, error) {
	return services.Func1(ctx, in)
}
