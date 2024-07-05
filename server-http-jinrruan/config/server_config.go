package config

import (
	"net/http"

	"git.code.oa.com/trpc-go/trpc-go/log"
)

var _serverHandler = &ServerHandler{server: http.DefaultServeMux, opts: &Options{}}

func init() {
	log.Info("do register")
	registerServerFilter()
	registerClientFilter()
}

type ServerHandler struct {
	server *http.ServeMux
	opts   *Options // 相当于服务的配置，里面包含filters
}

type Options struct {
	filters []ServerFilter
}

type Option func(options *Options)

var globalOptionLit = make([]Option, 0)

func loadOption() []Option {
	return globalOptionLit
}

func registerOption(opt ...Option) {
	globalOptionLit = append(globalOptionLit, opt...)
}

func WithFilter(filter ServerFilter) Option {
	return func(options *Options) {
		options.filters = append(options.filters, filter)
	}
}

func NewServerHandler(opt ...Option) *ServerHandler {
	optList := loadOption()
	optList = append(optList, opt...)
	for _, o := range optList {
		o(_serverHandler.opts)
	}
	return _serverHandler
}

func (s *ServerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, filter := range s.opts.filters {
		filter(writer, request)
	}
	s.server.ServeHTTP(writer, request)
}

type ServerFilter func(writer http.ResponseWriter, request *http.Request)

// 定义一个Handler接口实例，将http 默认Handler封装进来
func registerServerFilter() {
	registerOption(WithFilter(func(writer http.ResponseWriter, request *http.Request) {
		log.Info("====================【%+v】【%+v】【%+v】", request.Method, request.Host, request.RequestURI)
	}))
}

func registerClientFilter() {
}
