package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

func hash(key string) uint32 {
	hash := crc32.ChecksumIEEE([]byte(key))
	if key != "" {
		return hash ^ (hash >> 16)
	}
	return 0
}

func main() {
	key := "17740127052997474793"
	index := strconv.FormatUint(uint64(7&hash(key)), 10)
	fmt.Println("index: ", index)
}