package atomic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestBase(t *testing.T) {
	var a int64

	for i := 0; i < 100; i++ {
		atomic.AddInt64(&a, 2)
	}

	fmt.Println(a)
}
