package main

import (
	"errors"
	"fmt"
	"testing"
)

// defer 语句表达式的值在定义时就已经确定了
// 因为作为参数，err 在函数定义的时候就会求值，并且定义的时候 err 的值都是 nil，
// 所以最后打印的结果都是 nil;
func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("defer1 error")
	return
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("defer2 error")
	return
}

// 因为作为参数，err 在函数定义的时候就会求值，并且定义的时候 err 的值都是 nil，
// 所以最后打印的结果都是 nil;
func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer3 error")
	return
}

func e4() {
	var err error
	defer func(err *error) {
		fmt.Println(*err)
	}(&err)
	err = errors.New("defer4 error")
	return
}

func TestI6(*testing.T) {
	e1() //nil
	e2() //defer2 error
	e3() //nil
	e4() //defer4 error
}
