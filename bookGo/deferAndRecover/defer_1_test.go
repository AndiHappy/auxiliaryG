package main

import (
	"fmt"
	"testing"
)

// go 1.22.3 defer执行：2,1,0
// go 1.21 defer执行：2,2,2
// Go1.22之前版本for 循环声明的变量只创建一次，并在每次迭代中进行更新，这会导致遍历时访问value时实际上都是访问的同一个地址的值。
// Go1.22之后:
// • for循环的每次迭代都会定义新变量，而不再是共享一个变量
// • 支持对整数范围进行循环迭代
func TestDeferFormal(t *testing.T) {
	var whatever [3]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
