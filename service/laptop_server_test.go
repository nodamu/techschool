package service_test

import (
	"nodamu/pcbook/pb"
	"nodamu/pcbook/sample"
	"nodamu/pcbook/service"
	"testing"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
)


func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		laptop *pb.Laptop
		store service.LaptopStore
		code codes.Code
	} {
		{
			name: "success_with_id",
			laptop: sample.NewLaptop(),
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "success_no_id",
			laptop: sample.
		}
	}
}
