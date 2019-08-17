package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (dc DockerClient) ImageList() ([]types.ImageSummary, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ImageList(context.Background(), types.ImageListOptions{})
}

func (dc DockerClient) ImageRemove(imageID string) ([]types.ImageDeleteResponseItem, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ImageRemove(context.Background(), imageID, types.ImageRemoveOptions{})
}
