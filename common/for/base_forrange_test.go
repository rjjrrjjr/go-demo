package _for

import (
	"fmt"
	"git.code.oa.com/trpc-go/trpc-go"
	"log"
	"testing"
)

func TestForRange(t *testing.T) {

	stringChunk := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7"},
	}
	taskList := make([]func() error, 0)
	for _, curChunk := range stringChunk {
		curChunkTemp := curChunk
		taskList = append(taskList, func() error {
			fmt.Println(curChunkTemp)
			return nil
		})
	}

	err := trpc.GoAndWait(taskList...)
	if err != nil {
		log.Fatalln(err)
	}

}
