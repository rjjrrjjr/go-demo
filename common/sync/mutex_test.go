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

func async1(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("async1")
	time.Sleep(time.Second * 5)
}

func async2(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("async2")
	time.Sleep(time.Second * 5)
}

func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	go read("read1", cond)
	go read("read2", cond)
	go read("read3", cond)

	write("writer", cond)
	time.Sleep(time.Second * 2)

}

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()

	if !done {
		fmt.Println(name, "wait read")
		c.Wait()
	}

	fmt.Println(name, "start read")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "start writing")
	time.Sleep(time.Second * 2)
	done = true
	fmt.Println(name, "end writing")
	c.Broadcast()
}
