package client

import (
	"fmt"
	"github.com/docker/docker/client"
)

func (dc DockerClient) Deploy() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	fmt.Println(cli)
	return nil
}
