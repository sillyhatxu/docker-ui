package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	ctx context.Context
}

func New() *DockerClient {
	return &DockerClient{}
}

func (dc DockerClient) Info() (*types.Info, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	info, err := cli.Info(context.Background())
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (dc DockerClient) ClientVersion() (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	return cli.ClientVersion(), nil
}
