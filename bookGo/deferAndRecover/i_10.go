package main

import "fmt"

// 正常调用
func main() {
	defer h1()
	panic(404)
}
func h1() {
	if e := recover(); e != nil {
		fmt.Println("recover")
		return
	}
}

// 不能够调用 recover
func main0() {
	recover()
	panic(404)
}

// 不能够调用 recover
func main2() {
	defer recover()
	panic(404)
}

// 能。在 defer 的函数中调用recover，生效
func main3() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover")
		}
	}()
	panic(404)
}

// 能。在 defer 的函数中调用recover，生效
func main4() {
	defer func() {
		recover()
	}()
	panic(404)
}

// 不能。多重 defer 嵌套
func main5() {
	defer func() {
		defer func() {
			recover()
		}()
	}()
	panic(404)
}
