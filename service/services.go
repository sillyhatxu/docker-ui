package service

import (
	"fmt"
	"github.com/sillyhatxu/docker-ui/docker"
	"github.com/sillyhatxu/docker-ui/dto"
)

func ServiceList() ([]dto.ServiceDTO, error) {
	services, err := docker.Client.Services()
	if err != nil {
		return nil, err
	}
	var array []dto.ServiceDTO
	for _, service := range services {
		service.Endpoint.Ports[0].
		array = append(array, dto.ServiceDTO{
			Id:service.ID,
			Name:service.Spec.Name,
			Mode:fmt.Sprintf("%v", service.Spec.Mode.Replicated.Replicas),
			//Replicas:service.,
			Image:service.Spec.Labels["com.docker.stack.image"],
			Ports:service.Endpoint.Ports[0].
		})
	}
	if array == nil {
		array = make([]dto.ServiceDTO, 0)
	}
	return array, nil
}
