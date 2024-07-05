package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		file, fileHeader, _ := request.FormFile("file")
		fmt.Println("file header", fileHeader.Filename, fileHeader.Size)
		fileContent, _ := io.ReadAll(file)
		fmt.Println("file content: ", string(fileContent))
		fmt.Println("other arg, name: ", request.FormValue("name"))
	})
	_ = http.ListenAndServe(":10000", nil)
}
