package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	ctx context.Context
}

func PS() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ContainerList(context.Background(), types.ContainerListOptions{})
}
