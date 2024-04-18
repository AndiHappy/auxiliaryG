package gotutorial

import (
	"fmt"
	"testing"
)

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}

func TestDeferFunction(t *testing.T) {
	defer second()
	first()
}

func deferFunc() {
	fmt.Println("defer func called")
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	// 后执行
	defer deferFunc()
	// 先执行
	return returnFunc()
}

// 无命名返回值
func test() int {
	var i int
	defer func() {
		i++
		//作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=2
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		//作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=1
		fmt.Println("defer2", i)
	}()
	// 先执行return i, 把i的值给到一个临时变量，作为函数返回值
	return i
}

// 命名返回值的函数，由于返回值在函数定义的时候已经将该变量进行定义，
// 在执行return的时候会先执行返回值保存操作，
// 而后续的defer函数会改变这个返回值，虽然defer是在return之后执行的，
// 但是由于使用的函数定义的这个变量，所以执行defer操作后对该变量的修改，
// 进而最终会影响到函数返回值。
func test2() (i int) { //返回值命名i
	defer func() {
		i++
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2", i)
	}()
	return i
}

// 正常情况下，defer遇到return或者函数执行流程到达函数体末尾会将进入栈的defer出栈并以此执行，
// 同样遇到panic语句也是。遇到panic的时候，会遍历并将已经进栈的defer出栈并执行，
// 但是对于程序流程中panic之后的defer就不会进栈。在defer出栈执行的过程中，
// 遇到recover则停止panic，如果没有recover捕获panic，则执行完所以defer之后，抛出panic信息。
func test3() {
	defer func() { fmt.Println("defer: panic 之前0, 不捕获") }()
	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		// 捕获异常信息
		if err := recover(); err != nil {
			// 输出panic中的错误信息
			fmt.Println(err.(string))
		}
	}()
	// 正常进栈
	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()
	//触发defer出栈
	panic("触发异常")
	// 由于在panic之后，不会在执行
	defer func() {
		fmt.Println("defer: panic 之后, 永远执行不到")
	}()
}

func temperature(t float64) (rs string) {
	// write code here
	defer func() {
		if err := recover(); err != nil {
			rs = "体温异常"
		}
	}()

	if t > 37.5 {
		panic("体温异常")
	}

	return rs
}

func main1() {
	returnAndDefer()
	// defer 和 return之间的顺序是先返回值, i=0，后defer
	// 执行顺序为return语句->defer2->defer1->返回值。defer2先于defer1执行
	// return先执行，负责把结果写入返回值中，接着多个defer按照先进后出的顺序开始调用执行一些收尾工作，最后函数携带这个返回值退出。
	fmt.Println("test: ", test())

	//defer2 1
	//defer1 2
	//test: 2
	fmt.Println("test2:", test2())

	test3()
	// 由于存在recover捕获panic，main函数流程则正常执行
	fmt.Println("main 正常结束")

	fmt.Println(temperature(37))
	fmt.Println(temperature(40))
}
