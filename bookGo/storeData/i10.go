package main

import (
	"fmt"
	"time"
)

func chanTest1() {
	nonBlockChan := make(chan int)
	for true {
		select {
		case receive := <-nonBlockChan:
			fmt.Println("select receive: ", receive)
			break
		default:
			fmt.Println("select continue")
		}
	}
}

func chanTest2() {
	nonBlockChan := make(chan int)
	for {
		select {
		case nonBlockChan <- 1:
			println("sent 1 ")
		case c := <-nonBlockChan:
			println(" received ", c)
		}
		time.Sleep(time.Second)
	}
}

func chanTest3() {
	nonBlockChan := make(chan int, 1)
	for {
		select {
		case nonBlockChan <- 1:
			println("sent 1 ")
		case c := <-nonBlockChan:
			println(" received ", c)
		}
		time.Sleep(time.Second)
	}
}

// Go chan Invariants:
//
//	At least one of c.sendq and c.recvq is empty,
//	except for the case of an unbuffered channel with a single goroutine
//	blocked on it for both sending and receiving using a select statement,
//	in which case the length of c.sendq and c.recvq is limited only by the
//	size of the select statement.
func chanTest4() {
	var ch1 = make(chan int)
	go func() {
		goRoutine := 1
		for true {
			select {
			case ch1 <- 1:
				fmt.Println("sent 1") // just this runs
			case c := <-ch1:
				fmt.Println("received 1:", c)
			default:
				go func() {
					fmt.Println(goRoutine)
					goRoutine++
					time.Sleep(time.Second * 1000)
				}()
			}
		}
	}()
}

func main() {
	//chanTest2()
	//chanTest3()
	chanTest4()
	select {}
}
