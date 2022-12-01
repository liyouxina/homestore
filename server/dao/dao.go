package dao

import (
	"github.com/liyouxina/homestore/server/dao/client"
)

func InitDB() {
	client.Init()
}
