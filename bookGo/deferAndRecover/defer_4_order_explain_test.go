package main

import (
	"fmt"
	"testing"
)

// 重点是理解return t
// 并不是 一条原子性的语句
//
//	1)返回值 = xxx。
//	2)调用 defer 函数。
//	3)空的 return。
func f() (r int) {
	t := 1
	defer func() {
		t = t + 5
	}()
	fmt.Println("in method: ", t) //这个时候defer的方法还没有执行
	return t
}

func f_eaqual(r int) {
	t := 1
	// 1. 赋值指令
	r = t
	// 2. defer 被插入到赋值与返回之间执行，这个例子中返回值 r 没被修改过
	func() {
		t = t + 5
	}()
	// 3. 空的 return 指令
	return
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f3() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}

func TestI5(*testing.T) {
	fmt.Println(f())  //1
	fmt.Println(f2()) //1
	fmt.Println(f3()) //6
}
