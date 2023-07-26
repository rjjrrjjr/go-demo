package main

import "net/http"

type myHandler struct {
}

func (*myHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path{
	case "/1":
		rsp.Write([]byte("hehe"))
	case "/2":
		rsp.Write([]byte("haha"))
	}
}

func main() {

	// http.HandleFunc 会将pattern和对应的func注册到默认路由器中 DefaultServeMux = &defaultServeMux；m map[string]muxEntry
	// http.ListenAndServe("0.0.0.0:10000", nil) 不传handler，则会默认使用上面的 defaultServeMux
	//
	http.HandleFunc("/1", func(rsp http.ResponseWriter, req *http.Request) {
		rsp.Write([]byte("hello"))
	})
	http.HandleFunc("/2", func(rsp http.ResponseWriter, req *http.Request) {
		rsp.Write([]byte("world"))
	})

	http.ListenAndServe("0.0.0.0:10000", &myHandler{})
}