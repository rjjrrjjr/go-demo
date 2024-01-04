package nil

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {

	// 存在接口类型转换，则可能存在问题

	var a func()

	fmt.Println(a == nil)

	fmt.Println(interface{}(a) == nil)

	fmt.Println(getI() == nil)
}

func getI() interface{} {
	var a []uint64

	return a
}
