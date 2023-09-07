package gosdkexample

import (
	"fmt"
	"strings"
	"unsafe"
)

// ---------------------------------------------------compare 1-----------------------------------
// CompareT 这个方法非常的有意思，我们基本不使用这个方法，我们可以直接的比较，类似UserStringCompareExample的用法
// 直接比较了，就不需要调用方法了啊
func UserStringCompareExample() {
	if "aaa" > "aab" {
		fmt.Println("true")
	}

	if "aa0" < "aab" {
		fmt.Println("true")
	}

	if !("xxx" == "yyyy") {
		fmt.Println("false")
	}
}

func CompareT() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}

func CompareU(a, b string) int {
	// NOTE(rsc): This function does NOT call the runtime cmpstring function,
	// because we do not want to provide any performance justification for
	// using strings.Compare. Basically no one should use strings.Compare.
	// As the comment above says, it is here only for symmetry with package bytes.
	// If performance is important, the compiler should be changed to recognize
	// the pattern so that all code doing three-way comparisons, not just code
	// using strings.Compare, can benefit.
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

// ---------------------------------------------------compare 2-----------------------------------

// ---------------------------------------------------clone-----------------------------------
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

	// l := unsafe.IntegerType(len(b))
	// return unsafe.String(&b[0], len(b))

	return *(*string)(unsafe.Pointer(&b))
}

//---------------------------------------------------clone over -----------------------------------
