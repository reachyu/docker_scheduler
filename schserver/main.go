package main

import (
	"context"
	"log"
	"net"
	"schserver/apiserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "schserver/apis/go_docker"
)

const (
	port = ":8887"
)

// server is used to implement helloworld.GreeterServer.
type serverhello struct{}

// SayHello implements helloworld.GreeterServer
func (s *serverhello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &serverhello{})
	pb.RegisterContainerServiceServer(s, apiserver.NewDockerServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}