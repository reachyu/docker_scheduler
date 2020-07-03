package dockersvc

import (
	"context"
	"fmt"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerSvcInterface interface {
	CreateContainerInBackground(imageName string) string
	StopByContainerId(containerId string) error
	DeleteByContainerId(containerId string) error
	ListContainers()
}

type DockerSvc struct {
}

// 运行一个容器，相当于docker run alpine echo hello world
func createContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

// 后台运行容器，相当于键入 docker run -d bfirsh/reticulate-splines
func (s *DockerSvc) CreateContainerInBackground(imageName string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	//imageName := "bfirsh/reticulate-splines"

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
	return resp.ID
}

// 停止特定容器的日志
func (s *DockerSvc) StopByContainerId(containerId string) error{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err == nil {
		err = cli.ContainerStop(ctx, containerId, nil)
	}
	return err
}

// 删除特定容器的日志
func (s *DockerSvc) DeleteByContainerId(containerId string) error{
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err == nil {
		err = cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{})
	}
	return err
}

// 列出所有容器
func (s *DockerSvc) ListContainers() {
	//ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
		fmt.Println(container.Image)
		for lb := range container.Labels{
			fmt.Println(lb)
		}
		fmt.Println("=========================================================")
	}
}

// 打印特定容器的日志
func logsByContainerId(containerId string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, containerId, options)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}

// 列出并管理容器
func listAndManageContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
		// 停止容器
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			panic(err)
		}
	}
}

func NewDockerSvc() *DockerSvc {
	return &DockerSvc{}
}