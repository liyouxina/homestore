package client

import (
	"github.com/liyouxina/homestore/common/file_utils"
	"github.com/liyouxina/homestore/server/config"
	"github.com/liyouxina/homestore/server/dao/common"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type HomestoreClient struct {
	common.Common
	Name string `json:"name"`
	Path string `json:"path"`
}

func Insert(client *HomestoreClient) {
	db.Create(*client)
}

func QueryAll() []*HomestoreClient {
	var clients []*HomestoreClient
	db.Find(clients)
	return clients
}

func Init() {
	serverConfig, err := config.GetConfig()
	if err != nil {
		logrus.Error("init db error HomestoreClient", err)
		panic(err)
	}

	dbFullPath := serverConfig.DBPath + "/homestore.db"
	exists, _ := file_utils.PathExists(dbFullPath)
	if !exists {
		_ = file_utils.CreateFile(dbFullPath)
	}
	db, err := gorm.Open(sqlite.Open(dbFullPath))
	if err != nil {
		logrus.Error(err)
		logrus.Errorf("connect db error HomestoreClient %s", dbFullPath)
		panic(err)
	}
	if !exists {
		_ = db.AutoMigrate(&HomestoreClient{})
	}
}

