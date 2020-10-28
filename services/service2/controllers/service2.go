package controllers

// Load required packages
import (
	"context"

	model "github.com/kdelalic/go-microservice-structure/pkg/proto/service2"
	"github.com/kdelalic/go-microservice-structure/services/service2/services"
)

// Service2Server is the service2 server
type Service2Server struct {
}

// Func1 does something based on input
func (server *Service2Server) Func1(ctx context.Context, in *model.Func1Request) (*model.Func1Response, error) {
	return services.Func1(ctx, in)
}
