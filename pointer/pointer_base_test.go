package pointer

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangePointer(t *testing.T) {

	baseArr := make([]*string, 0)
	for i := 0; i < 1; i++ {
		a := strconv.FormatInt(int64(i), 10)
		b := strconv.FormatInt(int64(i), 32)
		c := strconv.FormatInt(int64(i), 36)
		baseArr = append(baseArr, &a, &b, &c)
	}
	fmt.Println(baseArr)

	resultMap := make(map[int]*string, 0)
	for i, s := range baseArr {
		resultMap[i] = s
	}
	fmt.Println(resultMap)
}

func TestRangePointer2(t *testing.T) {

	baseArr := make([]string, 0)
	for i := 0; i < 1; i++ {
		a := strconv.FormatInt(int64(i), 10)
		b := strconv.FormatInt(int64(i), 32)
		c := strconv.FormatInt(int64(i), 36)
		baseArr = append(baseArr, a, b, c)
	}
	fmt.Println(baseArr)

	resultMap := make(map[int]*string, 0)
	for i, s := range baseArr {
		temp := s
		resultMap[i] = &temp
	}
	fmt.Println(resultMap)
}

func TestStructWithPointerField(t *testing.T) {
	name := "name"
	age := 18
	a := A{
		Name: &name,
		Age: &age,
	}
	fmt.Println(*a.Name)

	*a.Name = "jinrruan"

	fmt.Println(*a.Name)
}

type A struct {
	Name *string
	Age *int
}