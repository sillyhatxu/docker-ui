package service

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceList(t *testing.T) {
	array, err := ServiceList()
	assert.Nil(t, err)
	for _, s := range array {
		logrus.Infof("%#v", s)
	}
}
