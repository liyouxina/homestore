package config

import (
	"errors"
	"flag"
	"github.com/liyouxina/homestore/common/string_utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	configFullPath = flag.String("config", "", "homestore server config file full path")
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
		logrus.Error("config file format error", err)
		return nil, err
	}

	isConfigValid, msg := checkConfigValid(config)
	if !isConfigValid {
		return nil, errors.New("config yaml is invalid " + msg)
	}

	return config, nil
}

type Config struct {
	Port int `yaml:"port"`
	ReceiverPort int `yaml:"receiver_port"`
	DBPath string `yaml:"db_path"`
	LogPath string `yaml:"log_path"`
}

func checkConfigValid(config *Config) (isValid bool, msg string) {
	if config == nil {
		return false, "config is nil"
	}
	if config.Port <= 0 || config.Port > 65535 {
		return false, "server port should more than 0 and less than 65535"
	}
	if config.ReceiverPort <= 0 || config.ReceiverPort > 65535 {
		return false, "receiver_port should more than 0 and less than 65535"
	}
	if config.DBPath == "" {
		return false, "db path is empty"
	}
	return true, ""
}