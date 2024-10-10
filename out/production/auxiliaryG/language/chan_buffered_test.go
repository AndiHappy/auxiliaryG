package gotutorial

import (
	"fmt"
	"testing"
	"time"
)

var done = make(chan bool)

func TestUNBlockChanRead(t *testing.T) {
	<-done
	fmt.Println("block forever")
}

func goroutineWrite() {
	fmt.Println("goroutineWrite start")
	fmt.Println("goroutineWrite over")
	done <- true
}

func TestUNBlockChanWrite(t *testing.T) {
	go goroutineWrite()
	<-done
	fmt.Println("test case over!")
	done <- true //永远阻塞在这里
	fmt.Println("block forever")
}

func TestUseSwitchForNonBlockChan(t *testing.T) {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func TestBroadcastMessage1(t *testing.T) {
	ch := make(chan string, 1)
	for i := 0; i < 3; i++ {
		go func() {
			msg := <-ch
			fmt.Println("Received message:", msg)
		}()
	}
	ch <- "Hello, World!"
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("test case over!")
}

func TestBroadcastMessage2(t *testing.T) {
	ch := make(chan string, 1)
	for i := 0; i < 3; i++ {
		go func() {
			msg := <-ch
			fmt.Println("Received message:", msg)
		}()
	}
	ch <- "Hello, World!"
	time.Sleep(time.Second)
	fmt.Println("test case over!")
}

func TestBroadcastMessage3(t *testing.T) {
	ch := make(chan string, 1)
	for i := 0; i < 3; i++ {
		go func() {
			if msg, ok := <-ch; ok {
				fmt.Println("Received message:", msg)
			}
		}()
	}
	ch <- "Hello, World!"
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("test case over!")
}
