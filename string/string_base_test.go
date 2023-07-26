package string

import (
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

func chunkArray(sourceArray []int, chunkSize int) [][]int {
	endIndex := chunkSize
	result := make([][]int, 0)
	for ; endIndex <= len(sourceArray); {
		result = append(result, sourceArray[endIndex - chunkSize : endIndex])
		endIndex = endIndex + chunkSize
	}
	if endIndex > len(sourceArray) {
		result = append(result, sourceArray[endIndex - chunkSize :])
	}
	return result
}

func TestSubString(t *testing.T) {
	numList := make([]int, 0)
	for i := 0; i < 1002; i++ {
		numList = append(numList, i)
	}

	fmt.Println(numList[:10])
	fmt.Println(numList[:10])

	// 500 一组打印
	result := chunkArray(numList, 200)
	for _, item := range result {
		fmt.Println(item)
	}
}

func TestI2A(t *testing.T) {
	fmt.Println("---------")
	a, b := strconv.ParseUint("32", 10, 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---------")
}

func TestStr(t *testing.T)  {
	a := "我是中国人"
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))
	fmt.Println(a[:1])
	fmt.Println(a[:2])
	fmt.Println(a[:3])
	fmt.Println(string([]rune(a)[:4]))
}

