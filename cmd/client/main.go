package main

import (
	"flag"
	"log"

	"github.com/nodamu/techschool/pb"
	"github.com/nodamu/techschool/sample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("create laptop with id: %s", res.Id)

}
