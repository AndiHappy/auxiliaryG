package main

import (
	"fmt"
	"testing"
)

func fib(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
func TestFib(t *testing.T) {
	ch := make(chan int) // create a buffered channel with a capacity of 10
	go fib(10, ch)       // generate the first 10 Fibonacci numbers in a separate goroutine
	// read values from the channel until it's closed
	for x := range ch {
		fmt.Println(x)
	}
}
