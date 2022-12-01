package main

import (
	"github.com/liyouxina/homestore/server"
	"github.com/liyouxina/homestore/server/config"
	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		err := recover()
		logrus.Error(err)
	}()
	serverConfig, err := config.GetConfig()
	if err != nil {
		logrus.Error("init server config error", err)
		panic(err)
	}
	instance := server.Builder().
		SetConfig(serverConfig).
		Build()
	instance.Run()
}
