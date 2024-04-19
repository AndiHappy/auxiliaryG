package gotutorial

import (
	"fmt"
	"runtime/debug"
	"testing"
	"time"
)

// 正确示例：
func T() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			//fmt.Println(string(debug.Stack()))
		}
	}()
	panic("Test(): panic")
}

func TestRecoverExample3(t *testing.T) {
	go T()
	time.Sleep(5 * time.Second)       // 模拟执行耗时任务(顺便等待子协程执行)
	fmt.Println("main()依然是能正常执行的...") // 可以正常打印，即使Test()发生panic
}
