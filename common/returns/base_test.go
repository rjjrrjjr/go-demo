package returns

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	doReturn()
}

func doReturn() (a int) {
	defer func() {
		fmt.Println(a)
	}()

	return 1000
}
