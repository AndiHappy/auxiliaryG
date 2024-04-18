package gotutorial

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestSelectNilChanNoDefault(t *testing.T) {
	var nilChan chan int
	select {
	case v := <-nilChan:
		fmt.Println(v)
	}
	//block
}

func TestSelectNilChan(t *testing.T) {
	var nilChan chan int
	select {
	case v := <-nilChan:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}
	//non block
}

func TestSelectNilChanForEver(t *testing.T) {
	var nilChan chan int
	for {
		select {
		case v := <-nilChan:
			fmt.Println(v)
		default:
			fmt.Println("default")
		}
	}
	//forever
}

func TestSelectChanTimeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 4)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		return
	}
}

func GetData(url string, ch chan int) {
	req, err := http.Get(url)
	if err != nil {

	} else {
		ch <- req.StatusCode
	}
}
func TestFastReturn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go GetData("https://www.2.com", ch1)
	go GetData("https://www.3.com", ch2)
	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	}
}

func TestWaitAndReturn(t *testing.T) {
	ch1 := make(chan int, 1)
	ch1 <- 1
	ch2 := make(chan int, 1)
	ch2 <- 2
	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println("信道1的结果：", v)
			case v := <-ch2:
				fmt.Println("信道2的结果：", v)
			default:
				continue
			}
		}
	}()
}

func TestSelectChanCloseEvent(t *testing.T) {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		close(c)
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func TestSelectChanExitFor(t *testing.T) {
	done := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
	}

	fmt.Printf("在通知退出循环时，执行了%d次.\n", workCounter)
}

func TestNilSelect(t *testing.T) {
	select {}
}
