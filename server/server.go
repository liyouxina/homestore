package server

import (
	"github.com/liyouxina/homestore/server/config"
)

type Server struct {
	config config.ServerConfig
}

func (s *Server) Run() {

}

func Builder() *serverBuilder {
	return new(serverBuilder)
}

type serverBuilder struct {
	server Server
}

func (s *serverBuilder) SetConfig(config config.ServerConfig) *serverBuilder {
	s.server.config = config
	return s
}

func (s *serverBuilder) Build() Server {
	return s.server
}
