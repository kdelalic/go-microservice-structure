package services

// Load required packages
import (
	"context"

	pb "github.com/kdelalic/go-microservice-structure/pkg/proto/service2"
)

// Func1 takes the input and returns some output
func Func1(ctx context.Context, req *pb.Func1Request) (*pb.Func1Response, error) {

	// Build response
	res := &pb.Func1Response{
		Rooms: nil,
	}

	return res, nil
}
