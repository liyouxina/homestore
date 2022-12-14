package server

import (
	"github.com/liyouxina/homestore/server/config"
	"github.com/liyouxina/homestore/server/dao"
	"github.com/liyouxina/homestore/server/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Server struct {
	config *config.Config
}

func Builder() *serverBuilder {
	return new(serverBuilder)
}

type serverBuilder struct {
	server Server
}

func (s *serverBuilder) SetConfig(config *config.Config) *serverBuilder {
	s.server.config = config
	return s
}

func (s *serverBuilder) Build() Server {
	return s.server
}

func (s *Server) Run() {
	dao.InitDB()
	httpHandler := handler.Handler{}
	err := http.ListenAndServe("127.0.0.1:" + strconv.Itoa(s.config.Port), httpHandler)
	if err != nil {
		logrus.Error(err)
		return
	}
}