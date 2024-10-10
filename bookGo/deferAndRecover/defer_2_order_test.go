package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// defer的执行顺序在return之前注册
func TestDeferOrder(t *testing.T) {
	defer func() {
		fmt.Println("before return")
	}()
	if true {
		fmt.Println("during return")
		return
	}
	defer func() {
		fmt.Println("after return")
	}()
}

// 典型的场景
func mergeFile() error {
	f, _ := os.Open("file1.txt")
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file1.txt err %v\n", err)
			}
		}(f)
	}

	// 打开文件二
	f, _ = os.Open("file2.txt")
	if f != nil { //首先进行f的nil判断
		defer func(f io.Closer) { //这里传入的f为打开的，即使后面被使用
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file2.txt err %v\n", err)
			}
		}(f)
	}
	return nil
}
