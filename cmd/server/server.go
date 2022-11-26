package main

import (
	"flag"
	"github.com/liyouxina/homestore/server"
	"github.com/liyouxina/homestore/server/config"
	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		err := recover()
		logrus.Error(err)
	}()
	instance := server.Builder().
		SetConfig(config.GetConfig()).
		Build()
	instance.Run()
}
