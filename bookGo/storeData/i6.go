package main

import (
	"fmt"
	"time"
)

func worker6() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}

// chan 作为执行定时执行周期性任务
func main() {
	go worker6()
	select {}
}
