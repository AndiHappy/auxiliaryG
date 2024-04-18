package main

import (
	"fmt"
	"time"
)

// chan作为接口生产者和消费者之间的媒介
func main() {
	taskCh := make(chan int, 100)
	go worker(taskCh)
	// 阻塞任务
	for i := 0; i < 10; i++ {
		taskCh <- i
	}
	// 等待 1 小时
	select {
	case <-time.After(time.Hour):
	}
}

func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
