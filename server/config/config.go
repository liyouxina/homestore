package config

import (
	"flag"
	"github.com/liyouxina/homestore/common/string_utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	port = flag.Int("port", 1999, "homestore server port")
)

var (
	defaultPath = []string{
		"./homestore.config",
		"/etc/homestore/homestore.config",
	}

	serverConfig *ServerConfig
)

func GetConfig(fullPath *string) (*ServerConfig, error) {
	if serverConfig != nil {
		return serverConfig, nil
	}
	if string_utils.IsEmpty(fullPath) {
		serverConfig = Builder().SetPort(*port).Build()
		return serverConfig, nil
	}
	serverConfig = &ServerConfig{}
	configFileContent, err := ioutil.ReadFile(*fullPath)
	if err != nil {
		logrus.Error("read config file error", err)
		return nil, err
	}
	err = yaml.Unmarshal(configFileContent, &serverConfig)
	if err != nil {
		logrus.Errorf("config file format error", err)
		return nil, err
	}
	return serverConfig, nil
}

type ServerConfig struct {
	Port int
}


type ServerConfigBuilder struct {
	serverConfig *ServerConfig
}

func Builder() *ServerConfigBuilder {
	return &ServerConfigBuilder{
		serverConfig: &ServerConfig{},
	}
}

func (s *ServerConfigBuilder) SetPort(port int) *ServerConfigBuilder {
	s.serverConfig.Port = port
	return s
}

func (s *ServerConfigBuilder) Build() *ServerConfig {
	return s.serverConfig
}