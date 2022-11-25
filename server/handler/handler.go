package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {

}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	request.URL.Path
}

func ()  {
	
}
