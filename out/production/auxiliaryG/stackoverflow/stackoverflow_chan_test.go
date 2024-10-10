package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/*
https://stackoverflow.com/questions/70089953/reading-from-a-goroutine-channel-without-blocking
*/

func helper(c chan<- error) {
	//do some work
	time.Sleep(5 * time.Second)
	c <- errors.New("") // send errors/nil on c
}

func worker() error {
	fmt.Println("do one")
	c := make(chan error, 1)
	go helper(c)

	err := <-c
	fmt.Println("do two")

	return err
}
func TestStackoverflow_chan(t *testing.T) {
	worker()
}
