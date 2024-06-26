package main

import (
	"fmt"
	"testing"
)

type Consumer struct {
	msgs *chan int
}

// NewConsumer creates a Consumer
func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

// consume reads the msgs channel
func (c *Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		msg := <-*c.msgs
		fmt.Println("consume: Received:", msg)
	}
}

// Producer definition
type Producer struct {
	msgs *chan int
	done *chan bool
}

// NewProducer creates a Producer
func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

// produce creates and sends the message through msgs channel
func (p *Producer) produce(max int) {
	fmt.Println("produce: Started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		*p.msgs <- i
	}
	*p.done <- true // signal when done
	fmt.Println("produce: Done")
}
func TestProducerConsumer(t *testing.T) {
	var msgs = make(chan int)  // channel to send messages
	var done = make(chan bool) // channel to control when production is done
	// Start a goroutine for Produce.produce
	go NewProducer(&msgs, &done).produce(5)
	// Start a goroutine for Consumer.consume
	go NewConsumer(&msgs).consume()
	// Finish the program when the production is done
	<-done
}
