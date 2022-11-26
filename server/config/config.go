package config

import (
	"flag"
	"github.com/liyouxina/homestore/common/string_utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	configFullPath = flag.String("config", "./homestore.yaml", "homestore server config file full path")
	port = flag.Int("port", 1999, "homestore server port")
	receiverPort = flag.Int("receiver-port", 2000, "receiver port")
	dbPath = flag.String("DBPath", "./", "db store path")
)

var (
	config *Config
)

func GetConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}
	if string_utils.IsEmpty(configFullPath) {
		config = &Config{
			Port: *port,
			DBPath: *dbPath,
			ReceiverPort: *receiverPort,
		}
		return config, nil
	}
	config = &Config{}
	configFileContent, err := ioutil.ReadFile(*configFullPath)
	if err != nil {
		logrus.Error("read config file error", err)
		return nil, err
	}
	err = yaml.Unmarshal(configFileContent, &config)
	if err != nil {
		logrus.Errorf("config file format error", err)
		return nil, err
	}
	return config, nil
}

type Config struct {
	Port int
	ReceiverPort int
	DBPath string
}