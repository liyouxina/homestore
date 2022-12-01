package config

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func TestGetConfig(t *testing.T) {

}

func TestIoutilReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("/this/is/a/file/not/existed")
	logrus.Error(data,err)
}