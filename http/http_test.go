package http

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			fmt.Println("do goroutine")
			time.Sleep(time.Second * 3)
			fmt.Println(writer.Write([]byte("hello world")))
			fmt.Println("------------")
		}()

		fmt.Println("do main")
		writer.Write([]byte("hello world"))
	})

	http.ListenAndServe(":10000", nil)
}

func TestLongPolling(t *testing.T) {

	resultChannel := make(chan string, 1)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			time.Sleep(time.Second * 2)
			resultChannel <- "DATA"
		}()

		select {
		case result := <-resultChannel:
			fmt.Fprintf(writer, result)
		case <-time.After(time.Second * 3):
			fmt.Fprintf(writer, "no change")
		}
	})
	http.HandleFunc("/timeout", func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			time.Sleep(time.Second * 5)
			resultChannel <- "DATA"
		}()

		select {
		case result := <-resultChannel:
			fmt.Fprintf(writer, result)
		case <-time.After(time.Second * 3):
			fmt.Fprintf(writer, "no change")
		}
	})

	http.ListenAndServe(":10000", nil)

}
