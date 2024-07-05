package _defer

import (
	"fmt"
	"testing"

	"git.code.oa.com/trpc-go/trpc-go/errs"
)

func TestDeferWithPanic(t *testing.T) {

	defer func() {
		fmt.Println("================defer")
	}()

	a := 0
	fmt.Println(1 / a)
	fmt.Println("----------------------end")
}

func TestA(t *testing.T) {
	user := &_user{}
	defer printUserName2(user.Name)
	user.Name = "jinrruan"
}

func TestB(t *testing.T) {
	_ = doB()
}

func doB() (err error) {
	defer doDefer(err)

	name, err := haha()

	_ = name

	return errs.New(8888888, "jinrruan")
}

func TestC(t *testing.T) {
	_ = doC()
}

func doC() (err error) {

	name, err := haha()

	defer func() {
		fmt.Println(err)
	}()

	_ = name

	return errs.New(8888888, "jinrruan")
}

func TestD(t *testing.T) {
	_ = doD()
}

func doD() (err error) {

	name, err := haha()
	_ = name

	defer doDefer(err)

	defer func() {
		doDefer(err)
	}()

	return errs.New(8888888, "jinrruan")
}

func haha() (string, error) {
	return "jnn", errs.New(99999999, "haha")
}

func doDefer(err error) {
	fmt.Println("----------")
	fmt.Println(err)
	fmt.Println("----------")
}

func printUserName(user *_user) {
	fmt.Println("hah", user.Name)
}

func printUserName2(name string) {
	fmt.Println("hah", name)
}

type _user struct {
	Name string
	Age  int
}
