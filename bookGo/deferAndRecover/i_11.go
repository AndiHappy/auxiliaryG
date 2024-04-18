package main

import (
	"fmt"
	"time"
)

var notifier chan interface{}

func startGlobalPanicCapturing() {
	notifier = make(chan interface{})
	go func() {
		for {
			select {
			case r := <-notifier:
				fmt.Println(r)
			}
		}
	}()
}
func main() {
	startGlobalPanicCapturing()
	// 产生恐慌，但该恐慌会被捕获
	Go(func() {
		a := make([]int, 1)
		println(a[1])
	})
	time.Sleep(time.Second * 10)
	fmt.Println("main over!")

}

// Go 是一个恐慌安全的 goroutine
func Go(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				notifier <- r
			}
		}()
		f()
	}()
}
