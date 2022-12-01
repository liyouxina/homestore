package receiver

import (
	"encoding/json"
	"errors"
	"github.com/liyouxina/homestore/common/go_utils"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

const SPLITER = "//////////"

const (
	METHOD_WHO = "who"
)

type Communicator struct {
	connection net.Conn
	sendMessagePool chan []byte
	receiveMessagePool chan []byte
	waitingPool map[int64]chan []byte
}

func (c *Communicator) RegisterInPool(pool map[string]*Communicator) error {
	id, err := c.GetClientName()
	if err != nil {
		logrus.Error("get", err)
		return err
	}
	pool[*id] = c
	return nil
}

func (c *Communicator) Keep() {
	defer func() {
		err := recover()
		if err != nil {
			logrus.Error("keep failed", err)
		}
		_ = c.connection.Close()
	}()
	c.keepSending()
	c.keepReceiving()
}

func (c *Communicator) GetClientName() (*string, error) {
	callResult, err := c.Call(METHOD_WHO, nil)
	if err != nil {
		logrus.Error("call who failed", err)
		return nil, err
	}
	id, exists, err := go_utils.GetStringFromMap("name", callResult)
	if !exists {
		logrus.Error("GetClientName no name")
		return nil, errors.New("GetClientName no name")
	}
	return id, nil
}

func (c *Communicator) keepSending() {
	for {
		sendMessage, _ := <- c.sendMessagePool
		_, err := c.connection.Write(sendMessage)
		if err != nil {
			logrus.Error("Communicator sendMessage error", err)
			panic(err)
		}
	}
}

func (c *Communicator) keepReceiving() {
	var store []byte
	spliterIndex := 0
	spliterSerious := []byte(SPLITER)
	for {
		buf := make([]byte, 1024)
		_, err := c.connection.Read(buf)
		if err != nil {
			logrus.Error("Communicator readMessage error", err)
			panic(err)
		}
		for bufIndex, bit := range buf {
			match := bit == spliterSerious[spliterIndex]
			if match {
				spliterIndex ++
			} else {
				spliterIndex = 0
			}
			if spliterIndex == len(spliterSerious) {
				spliterIndex = 0
				store = append(store, buf[0 : bufIndex]...)[0 : -len(spliterSerious)]
				c.receiveMessagePool <- store
			}
		}
	}
}

func (c *Communicator) Call(method string, params *string) (map[string]interface{}, error) {
	var message map[string]interface{}
	traceId := time.Now().UnixNano()
	message["traceId"] = traceId
	message["method"] = method
	message["params"] = params
	messageByte, _ := json.Marshal(message)
	c.sendMessagePool <- messageByte
	waitingChan := make(chan []byte, 1)
	c.waitingPool[traceId] = waitingChan
	data, _ := <- waitingChan

	return genBody(data)
}

func genBody(data []byte) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	err := json.Unmarshal(data, result)
	logrus.Infof("call return body %s", data)
	if err != nil {
		logrus.Error("call return error", err)
		return nil, err
	}
	isSuccessInterface, exists := result["success"]
	if !exists {
		logrus.Error("call result has no success")
		return nil, errors.New("call result has no success")
	}
	isSuccess := isSuccessInterface.(bool)
	if !isSuccess {
		logrus.Error("call result not success")
		return nil, errors.New("call result not success")
	}
	bodyInterface, exists := result["body"]
	if !exists {
		logrus.Error("call result has no body")
		return nil, errors.New("call result has no body")
	}
	body := map[string]interface{}{}
	err = json.Unmarshal([]byte(bodyInterface.(string)), body)
	if err != nil {
		logrus.Error("call body unmarshal failed")
		return nil, errors.New("call body unmarshal failed")
	}
	return body, nil
}