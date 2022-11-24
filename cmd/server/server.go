package main

import (
	"flag"
	"github.com/liyouxina/homestore/server"
	"github.com/liyouxina/homestore/server/config"
	"github.com/sirupsen/logrus"
)

var (
	configFullPath = flag.String("config", "", "homestore server config file full path")
	port           = flag.Int("port", 0, "homestore server port")
)

func main() {
	defer func() {
		err := recover()
		logrus.Error(err)
	}()
	serverConfig, err := config.GetConfig(configFullPath)
	if err != nil {
		logrus.Error(err)
	}
	instance := server.Builder().
		SetConfig(serverConfig).
		Build()
	instance.Run()
}
