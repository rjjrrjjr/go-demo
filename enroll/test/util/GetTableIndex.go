package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

func init() {
	fmt.Println("================asdfg")
}

func hash(key string) uint32 {
	hash := crc32.ChecksumIEEE([]byte(key))
	if key != "" {
		return hash ^ (hash >> 16)
	}
	return 0
}

func main() {
	key := "7864300907930846319"
	index := strconv.FormatUint(uint64(31&hash(key)), 10)
	fmt.Println("index: ", index)
}
