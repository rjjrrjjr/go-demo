package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexWithoutPointer(t *testing.T) {
	mutex := &sync.Mutex{}

	fmt.Println(1)
	go func() {
		async1(mutex)
	}()
	fmt.Println(2)
	go func() {
		async2(mutex)
	}()
	fmt.Println(3)
	time.Sleep(time.Second * 6)

}

func async1(mu *sync.Mutex)  {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("async1")
	time.Sleep(time.Second * 5)
}

func async2(mu *sync.Mutex)  {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("async2")
	time.Sleep(time.Second * 5)
}