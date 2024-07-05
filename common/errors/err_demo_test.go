package errs

import (
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {

	err := &BizErr{
		code:     10001,
		msg:      "haha",
		httpCode: 200,
	}

	code, msg := GetErrorCodeMessage(err)

	fmt.Println(code, msg)
}
