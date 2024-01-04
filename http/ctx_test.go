package http

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtxTimeOut3(t *testing.T) {
	fmt.Println(time.Now())
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("----------result", time.Now())
	})
	fmt.Println(time.Now())

	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())
	fmt.Println(timer.Stop())
}

func TestCtxTimeOut2(t *testing.T) {
	// 创建一个带有超时时间的 context 对象
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ctx2, cancel2 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel2()

	fmt.Println(time.Now())
	select {
	case <-ctx2.Done():
		// 超时时执行的逻辑
		fmt.Println("任务超时")
	case <-time.After(5 * time.Second):
		// 任务完成时执行的逻辑
		fmt.Println("任务完成")
	}
	fmt.Println(time.Now())
}

func TestCtxTimeOut(t *testing.T) {
	// 创建一个带有超时时间的 context 对象
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 在 goroutine 中执行任务
	go doTask(ctx)

	// 在 select 语句中等待任务完成或超时
	select {
	case <-ctx.Done():
		// 超时时执行的逻辑
		fmt.Println("任务超时")
	case <-time.After(5 * time.Second):
		// 任务完成时执行的逻辑
		fmt.Println("任务完成")
	}
}

func doTask(ctx context.Context) {
	fmt.Println("do task")
	// 模拟一个需要耗时的任务
	time.Sleep(2 * time.Second)

	// 检查是否需要提前退出任务
	select {
	case <-ctx.Done():
		// 超时或取消时提前退出任务
		fmt.Println("任务被取消或超时")
		return
	default:
		// 继续执行任务
		fmt.Println("任务继续执行")
		// ... 执行任务的逻辑
	}
}
