package array

import (
	"fmt"
	"strconv"
	"testing"
)

func TestForNullArray(t *testing.T) {
	var a []string

	for i, s := range a {
		fmt.Println(i, s)
	}
}

func TestLenArray(t *testing.T) {
	source := make([]string, 0)

	fmt.Println(len(source))
	result := chunkStringArray(source, 100)
	fmt.Println(len(result))
}

func TestChunkArr(t *testing.T) {
	sourceArray := make([]string, 0)
	for i := 0; i < 20; i++ {
		sourceArray = append(sourceArray, strconv.FormatInt(int64(i+1), 10))
	}
	fmt.Println(sourceArray)

	result := chunkStringArray(sourceArray, 50)
	fmt.Println(result)
}


func chunkStringArray(sourceArray []string, chunkSize int) [][]string {
	endIndex := chunkSize
	result := make([][]string, 0)
	if len(sourceArray) == 0 {
		return result
	}
	for endIndex <= len(sourceArray) {
		result = append(result, sourceArray[endIndex-chunkSize:endIndex])
		endIndex = endIndex + chunkSize
	}
	if endIndex-chunkSize < len(sourceArray) {
		result = append(result, sourceArray[endIndex-chunkSize:])
	}
	return result
}

func TestSlice(t *testing.T) {
	buf := make([]byte, 0, 4096) // big buffer
	buf = append(buf, []byte("abcd")...)

	_ = buf[4:]
}