package gotutorial

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func method1() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		fmt.Println("letter")
		for ch := 'A'; ch < 'Z'; ch += 1 {
			letter <- true
			fmt.Print(string(ch))
			<-number
		}
		close(letter)
	}()
	wg.Add(1)
	time.Sleep(time.Second)
	go func() {
		fmt.Println("number")
		start := 1
		for {
			number <- true
			fmt.Print(start)
			start += 1
			_, ok := <-letter
			if ok == false {
				break
			}
		}
		wg.Done()
	}()

	<-number
	wg.Wait()
	fmt.Print("\n")
}

func method2() {
	even := make(chan bool)
	odd := make(chan bool)
	go func() {
		defer close(odd)
		for i := 0; i <= 10; i += 2 {
			<-even
			print("Even ====>")
			println(i)
			odd <- true
		}
	}()
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		for i := 1; i <= 10; i += 2 {
			_, ok := <-odd
			if !ok {
				wait.Done()
				return
			}
			print("Odd ====>")
			println(i)
			even <- true
		}
	}()
	even <- true
	wait.Wait()
}

func TestAlternatelyPrint1(t *testing.T) {
	method2()
}
