package services

// Load required packages
import (
	"context"

	model "github.com/kdelalic/go-microservice-structure/pkg/proto/service2"
)

// Func1 takes the input and returns some output
func Func1(ctx context.Context, req *model.Func1Request) (*model.Func1Response, error) {

	// Build response
	res := &model.Func1Response{
		Rooms: nil,
	}

	return res, nil
}
