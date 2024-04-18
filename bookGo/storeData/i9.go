package main

import (
	"fmt"
	"time"
)

// chan 作为控制并发的goroutine的数目
func main() {
	go func() {
		var testChanIsNil chan int
		fmt.Println("testChanIsNil")
		intv := <-testChanIsNil
		fmt.Println(intv)
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("main over!")

}
