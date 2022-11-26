package handler

import (
	"fmt"
	"github.com/liyouxina/homestore/server/dao/client"
	"net/http"
)

const (
	HEAD_TPL = `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN"><html>
<title>Directory listing for /</title>
<body>
<h2>Directory listing for /</h2>
<hr>
<ul>`
	TAIL_TPL = `</ul>
<hr>
</body>
</html>
`
	ITEM_TPL = `<li><a href="%s">%s</a>`
)

type Handler struct {

}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte(HEAD_TPL))
	path := request.URL.Path
	switch path {
	case "/": {
		handleRoot(writer)
		break
	}
	case "": {
		handleRoot(writer)
		break
	}
	default:



	}
	if path == "/" {

	}
	writer.Write([]byte(TAIL_TPL))
}

func handleRoot(writer http.ResponseWriter) {
	clients := client.QueryAll()
	for _, client := range clients {
		itemContent := client.Path + "/"
		writer.Write([]byte(fmt.Sprintf(ITEM_TPL, itemContent)))
	}
}

func handleNormal(writer http.ResponseWriter)  {

}
