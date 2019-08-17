package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (dc DockerClient) Containers() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ContainerList(context.Background(), types.ContainerListOptions{})
}

func (dc DockerClient) ContainerStats(containerID string) (*types.ContainerStats, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	containerStats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return nil, err
	}
	return &containerStats, nil
}
