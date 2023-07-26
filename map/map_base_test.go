package _map

import (
	"fmt"
	"testing"
)

func TestListMap(t *testing.T) {

	baseMap := make(map[string][]string, 0)

	str := "abc"

	baseMap[str] = make([]string, 0)
	baseList := baseMap[str]

	addElem(&baseList)
	fmt.Println(baseList)

	fmt.Println(baseMap)
}

func addElem(baseListPointer *[]string) {
	str := "abc"
	*baseListPointer = append(*baseListPointer, str)
	*baseListPointer = append(*baseListPointer, str)
	*baseListPointer = append(*baseListPointer, str)
}

func TestAddList(t *testing.T) {
	a := make([]string, 0)
	a = append(a, "123")
	addAnotherElement(&a)
	fmt.Println(a)
}

func addAnotherElement(a *[]string) {
	*a = append(*a, "456")
}

func TestMapKey2Array(t *testing.T) {
	baseMap := make(map[int]struct{}, 0)

	for i := 0; i < 10; i++ {
		baseMap[i] = struct{}{}
	}
	fmt.Println(baseMap)

	result := make([]int, 0)
	for key := range baseMap {
		result = append(result, key)
	}

	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	baseMap := make(map[string]int, 0)
	baseMap["1"] = 1
	baseMap["2"] = 2
	baseMap["3"] = 3

	a := &A{
		baseMap: baseMap,
	}
	fmt.Println(a)
	delete(a.baseMap, "2")
	fmt.Println(a)
}

type A struct {
	baseMap map[string]int
}

func TestDelInRange(t *testing.T) {
	baseMap := make(map[string]int, 0)
	baseMap["1"] = 1
	baseMap["2"] = 2
	baseMap["3"] = 3

	for key, _ := range baseMap {
		fmt.Println(key)
		delete(baseMap, "2")
	}

	fmt.Println(baseMap)
}