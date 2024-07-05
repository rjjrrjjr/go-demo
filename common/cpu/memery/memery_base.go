package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func test() {
	var s, e runtime.MemStats
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)
	time.Sleep(5 * time.Second) //睡眠，保持程序不退出
	log.Println(" ===> loop begin.")
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println(" ===> loop end.")
	runtime.ReadMemStats(&e)
	fmt.Println(e.Alloc, bToMb(e.Alloc))
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

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
