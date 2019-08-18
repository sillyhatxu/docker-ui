package client

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io/ioutil"
)

func (dc DockerClient) Containers() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli.ContainerList(context.Background(), types.ContainerListOptions{})
}

func (dc DockerClient) containerStats(containerID string) (*types.ContainerStats, error) {
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

func (dc DockerClient) ContainerStats(containerID string) (*ContainerStats, error) {
	containerStats, err := dc.containerStats(containerID)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(containerStats.Body)
	if err != nil {
		return nil, err
	}
	cs := ContainerStats{OSType: containerStats.OSType}
	var cssb StatsBody
	err = json.Unmarshal(body, &cssb)
	if err != nil {
		return nil, err
	}
	cs.Body = &cssb
	return &cs, nil
}
