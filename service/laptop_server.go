package service

import (
	"context"
	"errors"
	"github.com/nodamu/techschool/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides the laptop service
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopServer returns a new NewLaptopServer
func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recieve a create-laptop request with id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop id: %v", err)
		}
		laptop.Id = id.String()
	}
	//save the laptop to in-memory store
	err := server.Store.Save(laptop)

	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrorAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)

	}

	log.Printf("Saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
