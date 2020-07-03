package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "schclient/apis/go_docker"
)

const (
	address     = "localhost:8887"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	testlistimages(conn)
	//testhello(conn)
}

func testlistimages(conn *grpc.ClientConn)  {
	c := pb.NewContainerServiceClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.ListContainers(ctx,&empty.Empty{})
}

func testhello(conn *grpc.ClientConn)  {
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}