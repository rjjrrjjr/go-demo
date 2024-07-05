package config

import (
	"net/http"

	"git.code.oa.com/trpc-go/trpc-go/log"
)

var _serverHandler = &ServerHandler{server: http.DefaultServeMux}

func init() {
	log.Info("do register")
	registerServerFilter()
	registerClientFilter()
}

func NewServerHandler() *ServerHandler {
	return _serverHandler
}

type ServerHandler struct {
	server  *http.ServeMux
	filters []ServerFilter
}

func (s *ServerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, filter := range s.filters {
		filter(writer, request)
	}
	s.server.ServeHTTP(writer, request)
}

type ServerFilter func(writer http.ResponseWriter, request *http.Request)

// 定义一个Handler接口实例，将http 默认Handler封装进来
func registerServerFilter() {
	_serverHandler.filters = append(_serverHandler.filters, func(writer http.ResponseWriter, request *http.Request) {
		log.Info("====================【%+v】【%+v】【%+v】", request.Method, request.Host, request.RequestURI)
	})
}

func registerClientFilter() {
}
