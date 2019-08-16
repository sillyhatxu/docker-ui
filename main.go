package main

import (
	"flag"
	"fmt"
	"github.com/sillyhatxu/docker-ui/api"
	"github.com/sillyhatxu/docker-ui/config"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	flag.StringVar(&config.Conf.ServerHost, "p", "0.0.0.0:8080", "local server address")
	flag.Usage = func() {
		fmt.Println(fmt.Sprintf("Usage of %s:\n", os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	err := api.InitialAPI()
	if err != nil {
		logrus.Errorf("elasticsearch-ui error. Error : %v", err)
	} else {
		logrus.Info("project close")
	}
}
