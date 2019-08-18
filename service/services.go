package service

import (
	"fmt"
	"github.com/sillyhatxu/docker-ui/docker"
	"github.com/sillyhatxu/docker-ui/dto"
	"strconv"
)

func ServiceList() ([]dto.ServiceDTO, error) {
	services, err := docker.Client.Services()
	if err != nil {
		return nil, err
	}
	var array []dto.ServiceDTO
	for _, service := range services {
		//logrus.Infof("%#v",service)
		var ports string
		configArray := service.Endpoint.Ports
		for i, config := range configArray {
			if i > 0 {
				ports += fmt.Sprintf(",*:%s->%s/%s", fmt.Sprint(config.PublishedPort), fmt.Sprint(config.TargetPort), config.Protocol)
			} else {
				ports += fmt.Sprintf("*:%s->%s/%s", fmt.Sprint(config.PublishedPort), fmt.Sprint(config.TargetPort), config.Protocol)
			}
		}
		//*:9200->9200/tcp, *:9300->9300/tcp
		array = append(array, dto.ServiceDTO{
			Id:       service.ID,
			Name:     service.Spec.Name,
			Mode:     "replicated",
			Replicas: strconv.FormatUint(*service.Spec.Mode.Replicated.Replicas, 10),
			Image:    service.Spec.Labels["com.docker.stack.image"],
			Ports:    ports,
		})
	}
	if array == nil {
		array = make([]dto.ServiceDTO, 0)
	}
	return array, nil
}
