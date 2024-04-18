package main

import (
	"fmt"
	"time"
)

var token = make(chan int, 3)

// chan 作为控制并发的goroutine的数目
func main() {
	work := make([]func(int), 10)
	for i := 0; i < 10; i++ {
		work[i] = func(i int) {
			fmt.Println("method:", i)
			time.Sleep(time.Second)
		}
	}
	for i, w := range work {
		var ii = i
		go func() {
			token <- 1
			w(ii)
			<-token
		}()
	}

	// 等待 1 小时
	select {
	case <-time.After(time.Hour):
	}
}
