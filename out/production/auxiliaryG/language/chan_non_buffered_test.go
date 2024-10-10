package gotutorial

import (
	"fmt"
	"github.com/AndiHappy/auxiliaryG/util"
	"testing"
)

var doneL = make(chan bool, 3)

func blockGoroutineWrite() {
	fmt.Println("blockGoroutineWrite start")
	fmt.Println("blockGoroutineWrite over")
	doneL <- true
}

func TestBlockChanWrite(t *testing.T) {
	go blockGoroutineWrite()
	<-doneL
	fmt.Println("test case over!")
	doneL <- true //不会阻塞在这里
	fmt.Println("nonblock")
}

func TestUNBlockChanSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		//close(ch1) 如果不关闭的话，"test case over!" 就不会输出
	}()
	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		//close(ch2)
	}()
	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", x)
			} else {
				fmt.Println("ch1 closed")
			}
		case x, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", x)
			} else {
				fmt.Println("ch2 closed")
			}
		}
	}

	fmt.Println("test case over!")
}

func TestReceiveFromClosedChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received", i)
	}
}

func TestChanAsPara(t *testing.T) {
	v := make(chan int)
	go readData(v)
	sendData(v)
}
func readData(ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received", i)
	}
}
func sendData(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func TestChanIsFul(t *testing.T) {
	channel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Println(util.IsFull(channel))
	}
}
