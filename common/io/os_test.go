package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestStartProcess(t *testing.T) {

	f, _ := os.Open("a.txt")
	defer f.Close()

	attr := &os.ProcAttr{
		Files: []*os.File{f, os.Stdout, os.Stderr},
	}

	p, _ := os.StartProcess("/bin/cat", []string{"/"}, attr)

	p.Wait()

}

func TestIOReadALl(t *testing.T) {
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * 2000)
			w.WriteString(strconv.FormatInt(int64(i), 10) + "\n")
			w.WriteString(io.EOF.Error())
			w.WriteString("\n")
			fmt.Println("write success ", i)
		}
	}()

	for {
		fmt.Println("ready to read")
		// io.ReadAll() 函数会一直等待，直到遇到 EOF 标志或者管道被关闭
		content, _ := io.ReadAll(r)
		fmt.Printf("read: %s\n", content)
	}
}

func TestScanner(t *testing.T) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	fmt.Println("end")

}

func TestPipe(t *testing.T) {
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * 100)
			w.WriteString(strconv.FormatInt(int64(i), 10) + "\n")
			fmt.Println("write success ", i)
		}
	}()

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			str := s.Text()
			fmt.Println("read: ", str)
			if str == "end" {
				break
			}
		}
		fmt.Println("=======================================read end")
	}()
	time.Sleep(time.Hour)
	//io.Copy(os.Stdout, r)
}
