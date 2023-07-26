package _defer

import (
	"fmt"
	"testing"
)

func TestDeferWithPanic(t *testing.T) {

	defer func() {
		fmt.Println("================defer")
	}()

	a := 0
	fmt.Println(1/a)
	fmt.Println("----------------------end")
}
