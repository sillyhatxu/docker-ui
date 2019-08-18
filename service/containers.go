package service

import (
	"github.com/sillyhatxu/docker-ui/docker"
	"github.com/sillyhatxu/docker-ui/dto"
)

func DockerStats() ([]dto.DockerStatsDTO, error) {
	containers, err := docker.Client.Containers()
	if err != nil {
		return nil, err
	}
	var array []dto.DockerStatsDTO
	for _, container := range containers {
		containerStats, err := docker.Client.ContainerStats(container.ID)
		if err != nil {
			return nil, err
		}
		array = append(array, dto.DockerStatsDTO{
			ContainerId:   containerStats.GetId(),
			Name:          containerStats.GetName(),
			Cpu:           containerStats.GetCPUPercent(),
			MemUsageLimit: "",
			Mem:           "",
			NetIO:         "",
			BlockIO:       "",
			Pids:          containerStats.GetPids(),
		})
	}
	if array == nil {
		array = make([]dto.DockerStatsDTO, 0)
	}
	return array, nil
}
