package receiver

import (
	"github.com/liyouxina/homestore/server/config"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

type Receiver struct {
	pool map[string]*Communicator
}

func (r *Receiver) Run() {
	serverConfig, _ := config.GetConfig()
	listen, err := net.Listen("tcp", "127.0.0.1:" + strconv.Itoa(serverConfig.ReceiverPort))
	if err != nil {
		logrus.Error("run receiver error", err)
		panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			logrus.Error("get connection error", err)
			continue
		}
		communicator := &Communicator{
			connection: conn,
		}
		err = communicator.RegisterInPool(r.pool)
		if err != nil {
			logrus.Error("communicator RegisterInPool error", err)
			continue
		}
		communicator.Keep()
	}
}

func (r *Receiver) Get()  {
	
}