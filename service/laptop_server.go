package service

import (
	"context"
	"log"
	"nodamu/pcbook/pb"
)

// LaptopServer is the server that provides the laptop service
type LaptopServer struct {
}

// NewLaptopServer returns a new NewLaptopServer
func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recieve a create-laptop request with id: %s", laptop.Id)

	if

}
