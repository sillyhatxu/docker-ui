package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"io"
)

func (dc DockerClient) Services() ([]swarm.Service, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ServiceList(context.Background(), types.ServiceListOptions{})
}

func (dc DockerClient) ServiceRemove(serviceID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	return cli.ServiceRemove(context.Background(), serviceID)
}

func (dc DockerClient) ServiceLogs(serviceID string) (io.ReadCloser, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ServiceLogs(context.Background(), serviceID, types.ContainerLogsOptions{})
}
