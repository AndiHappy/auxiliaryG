package gosdkexample

import (
	"fmt"
	"strings"
	"unsafe"
)

func CloneT() {
	s := "hello,kk0man"
	clone := strings.Clone(s)
	fmt.Println(clone)

	// clone 可以使用 == 进行判断
	fmt.Println("s == clone: ", s == clone)

	//但是起始的初始地址是不相同
	fmt.Println("unsafe.StringData(s) == unsafe.StringData(clone): ", unsafe.StringData(s) == unsafe.StringData(clone))

	sByte := unsafe.StringData(s)
	fmt.Println("%V", sByte)

	cloneByte := unsafe.StringData(clone)
	fmt.Println("%V", cloneByte)

	fmt.Println("&s:=", &s)
	fmt.Println("&clone=", &clone)
}

func CloneS(s string) string {
	if len(s) == 0 {
		return ""
	}
	b := make([]byte, len(s))
	copy(b, s)
	// 其中的逻辑是:
	//	&b[0] 开始的位置，
	//  unsafe.IntegerType(len(b))  字符串的长度

	//l := unsafe.IntegerType(len(b))

	return *(*string)(unsafe.Pointer(&b))
}
