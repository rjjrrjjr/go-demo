package handler

import (
	"net/http"
	"sync"
)

var _baseHandler *baseHandler
var once = &sync.Once{}

func GetBaseHandler() *baseHandler {
	if _baseHandler != nil {
		return _baseHandler
	}
	once.Do(func() {
		_baseHandler = &baseHandler{}
	})
	return _baseHandler
}

type baseHandler struct {
}

func (b *baseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello base"))
}
