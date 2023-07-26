package io

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {


	file, err := os.OpenFile("a.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, fs.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, _ = file.WriteString("1234\n")
	_, _ = file.WriteString("1234\n")
	_, _ = file.WriteString("1234\n")
}

func TestCreateFile2(t *testing.T) {
	err := ioutil.WriteFile("b.txt", []byte("123456"), fs.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("文件已成功写入.")
}

func TestCreateFile3(t *testing.T) {
	file, err := os.Create("c.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("123456")
	if err != nil {
		log.Println(err)
	}

	log.Println("文件已成功写入.")
}