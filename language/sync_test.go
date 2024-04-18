package gotutorial

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Sync(t *testing.T) {
	// sync.Once 借助 m Mutex 实现
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		var ii = i
		go func() {
			once.Do(onceBody)
			fmt.Println("goroutine finished successfully ", ii)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
