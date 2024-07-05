package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// 自定义base64
const selfEncodeStd = "ABCDabcdefghijklEFGHIJKLMNOPQRSTUVWXYZmnopqrs56789tuvwxyz01234+/"

// EncodeSelf 自定义base64加密
func EncodeSelf(target []byte) string {
	return base64.NewEncoding(selfEncodeStd).EncodeToString(target)
}

// DecodeSelf 自定义base64解密
func DecodeSelf(target string) ([]byte, error) {
	return base64.NewEncoding(selfEncodeStd).DecodeString(target)
}

func TestDecodeEnrollAnswer(t *testing.T) {

	answerDetailList := []string{"kDa8jDIwiuE9Edb9hmj7PE==", "iGvviHexiXIyjDa8"}

	for _, answerDetail := range answerDetailList {
		result, err := DecodeSelf(answerDetail)
		fmt.Println(string(result), err)
	}

}
