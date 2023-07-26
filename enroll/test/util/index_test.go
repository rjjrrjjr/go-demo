package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestGetRegisterRecordIndex(t *testing.T) {
	fmt.Println("----------------")
	eventId := "E167607343331775283227129"
	md5Sum := md5.Sum([]byte(eventId))
	hashHex := hex.EncodeToString(md5Sum[14:16])
	if hashNum, err := strconv.ParseUint(hashHex, 16, 64); err != nil {
		fmt.Println(-1)
	} else {
		fmt.Println(int(hashNum) % 32)
	}
	fmt.Println("=================")
}

func TestIndex2(t *testing.T) {
	fmt.Println("----------------")
	shardValue := "E167353265340755148827129"
	partitionCount := 32
	md5Sum := md5.Sum([]byte(shardValue))
	hashHex := hex.EncodeToString(md5Sum[14:16])
	if hashNum, err := strconv.ParseUint(hashHex, 16, 64); err != nil {
		fmt.Println(-1)
	} else {
		fmt.Println(int(hashNum) % partitionCount)
	}
	fmt.Println("=================")
}

// user_ticket_record 获取index，根据uid获取
func TestIndex3(t *testing.T) {
	shardValue := "jinrruan"
	if _, err := strconv.ParseInt(shardValue, 10, 64); err != nil {
		strategyId := ShardMd5(shardValue, 32)
		fmt.Println(strategyId)
	} else {
		strategyId := ShardByLast5Char(shardValue, 32)
		fmt.Println(strategyId)
	}
}

// user_ticket_stub 获取index，根据ticket_record_id获取
func TestIndex4(t *testing.T) {
	shardValue := "T168113985042975129633077"
	fmt.Println(ShardByLast5Char(shardValue, 32))
}

// ShardMd5 字符串md5后取最后4位16进制数转10进制(最大65536)后取模，
func ShardMd5(shardValue string, partitionCount int) int {
	md5Sum := md5.Sum([]byte(shardValue))
	hashHex := hex.EncodeToString(md5Sum[14:16])
	if hashNum, err := strconv.ParseUint(hashHex, 16, 64); err != nil {
		return -1
	} else {
		return int(hashNum) % partitionCount
	}
}

func ShardByLast5Char(shardValue string, partitionCount int) int {
	if len(shardValue) >= 5 {
		if v, err := strconv.Atoi(shardValue[len(shardValue)-5:]); err != nil {
			return -1
		} else {
			return v % partitionCount
		}
	}
	return -1
}