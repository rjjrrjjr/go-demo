package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func test() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)
	time.Sleep(20 * time.Second) //睡眠，保持程序不退出
	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println(" ===> loop end.")
}

func main() {
	log.Println("Start.")

	//启动pprof
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

	test()

	log.Println("force gc.")
	runtime.GC() //强制调用gc回收

	log.Println("Done.")

	time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}