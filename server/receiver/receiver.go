package receiver

import (
	"bufio"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/liyouxina/homestore/server/config"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

var ReceiverPool map[string]*Receiver

type Receiver struct {
	pool map[string]*Communicator
}

func (r *Receiver) GetPathContent(path string) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(r.connection)
	writer.
}



func Run()  {
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
		communicator.Keep()
	}
}

func handleReceiver(conn net.Conn) {
	receiver := Receiver{
		connection: conn,

	}
	reader := bufio.NewReader(conn)
	reader.
	conn.Read()
	conn.Write()
}
