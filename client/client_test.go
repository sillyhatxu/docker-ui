package client

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var Client = New()

func TestContainers(t *testing.T) {
	containers, err := Client.Containers()
	assert.Nil(t, err)
	for _, container := range containers {
		logrus.Infof("%#v", container)
	}
}

//func TestContainerStats(t *testing.T) {
//	containers, err := Client.Containers()
//	assert.Nil(t, err)
//	for _, container := range containers {
//		containerStats, err := Client.ContainerStats(container.ID)
//		assert.Nil(t, err)
//		body, err := ioutil.ReadAll(containerStats.Body)
//		assert.Nil(t, err)
//		logrus.Infof("%s | %s", containerStats.OSType, string(body))
//	}
//}

func TestContainerStats(t *testing.T) {
	containers, err := Client.Containers()
	assert.Nil(t, err)
	for _, container := range containers {
		containerStats, err := Client.ContainerStats(container.ID)
		assert.Nil(t, err)
		logrus.Infof("%v,%v,%v,%v", containerStats.GetId(), containerStats.GetName(), containerStats.GetPids(), containerStats.GetCPUPercent())
	}
}

func TestServices(t *testing.T) {
	services, err := Client.Services()
	assert.Nil(t, err)
	for _, service := range services {
		logrus.Infof("%#v", service)
	}
}
