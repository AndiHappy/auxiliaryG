package gotutorial

import (
	"fmt"
	"testing"
	"time"
)

func worker(input chan int, output chan int, done chan bool) {
	for {
		select { //Use a select statement to receive values from multiple channels
		case n := <-input:
			// Do some work on n
			output <- n * 2
		case <-done:
			close(output) //Only the sender should close the channel
			return
		}
	}
}
func TestMulChanReceive(t *testing.T) {
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool)
	go worker(input, output, done)
	// Send some values on the input channel
	for i := 0; i < 10; i++ {
		input <- i
	}
	// Close the input channel to signal the end of input
	close(input)
	// Receive values from the output channel
	for n := range output {
		fmt.Println(n)
	}
	// Signal the worker to exit
	done <- true
}

func TestRangeOverChannelsVSSelect1(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for val := range ch {
		fmt.Println(val)
	}
}

func TestRangeOverChannelsVSSelect2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 10)
		ch2 <- 2
	}()
	for i := 0; i < 2; i++ {
		select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
		case val := <-ch2:
			fmt.Println("Received from ch2:", val)
		}
	}
}
