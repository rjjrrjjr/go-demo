package main

import (
	"net/http"

	"git.code.oa.com/trpc-go/trpc-go/log"
	"jinrruan-demo/server-http-jinrruan/config"
	"jinrruan-demo/server-http-jinrruan/handler"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello jinrruan"))
	})

	http.Handle("/hello", handler.GetBaseHandler())

	http.ListenAndServe(":8080", config.NewServerHandler())

	log.Info("==================start success")
}
