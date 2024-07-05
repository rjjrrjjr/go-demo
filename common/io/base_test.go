package io

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"jinrruan-demo/common/io/inner"
)

func TestCreateFile(t *testing.T) {

	file, err := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, fs.ModePerm)
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

func TestPath(t *testing.T) {
	dir := "."
	name := "/a.html"
	fmt.Println(filepath.Join(dir, name))
	inner.RequestOuterFile()
}

func TestBufio(t *testing.T) {
	str := "Hello, World!"
	reader := strings.NewReader(str)
	bufReader := bufio.NewReader(reader)

	// 读取一行
	line, _ := bufReader.ReadString(',')

	fmt.Println(line) // 输出: Hello,

	// 重新设置 bufio.Reader 关联的底层读取流
	newReader := strings.NewReader("12345")
	bufReader.Reset(newReader)

	// 读取切换后的底层读取流
	line, _ = bufReader.ReadString('\n')

	fmt.Println(line) // 输出: 12345
}

func TestBufio2(t *testing.T) {
	var wh writeHook
	defaultWriter := bufio.NewWriterSize(&wh, 200)
	defaultWriter.WriteByte(0)
	defaultWriter.Flush()
	fmt.Println(defaultWriter.Size())
	fmt.Println(cap(wh.p))
}

type writeHook struct {
	p []byte
}

func (wh *writeHook) Write(p []byte) (int, error) {
	wh.p = p
	return len(p), nil
}
