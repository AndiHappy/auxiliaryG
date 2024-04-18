package gotutorial

import (
	"fmt"
	"runtime/debug"
	"time"
)

// 正确示例：
func Test() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			//fmt.Println(string(debug.Stack()))
		}
	}()
	panic("Test(): panic")
}

func main3() {
	go Test()
	time.Sleep(5 * time.Second)       // 模拟执行耗时任务(顺便等待子协程执行)
	fmt.Println("main()依然是能正常执行的...") // 可以正常打印，即使Test()发生panic
}
