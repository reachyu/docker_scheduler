package apiserver

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "schserver/apis/go_docker"
	"schserver/apiserver/server"
)

type DockerServer struct {
}

func (s *DockerServer) RunContainer(ctx context.Context, request *api.ReqImageInfo) (*api.RunResponse, error) {
	smn := server.NewSvcManager()
	containerId := smn.CreateContainerInBackground(request.ImageName)
	return &api.RunResponse{Message: string(containerId)}, nil
}

func (s *DockerServer) StopContainer(ctx context.Context, request *api.ReqContainerInfo) (*api.StopResponse, error) {
	smn := server.NewSvcManager()
	err := smn.StopByContainerId(request.ContainerId)
	return &api.StopResponse{Message: string(err.Error())}, nil
}

func (s *DockerServer) DeleteContainer(ctx context.Context, request *api.ReqContainerInfo) (*api.DeleteResponse, error) {
	smn := server.NewSvcManager()
	err := smn.DeleteByContainerId(request.ContainerId)
	return &api.DeleteResponse{Message: string(err.Error())}, nil
}

func (s *DockerServer) ListContainers(ctx context.Context, request *empty.Empty) (*empty.Empty, error) {
	smn := server.NewSvcManager()
	smn.ListContainers()
	return &empty.Empty{}, nil
}

func NewDockerServer() *DockerServer {
	return &DockerServer{}
}