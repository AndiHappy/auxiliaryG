package gosdkexample

import (
	"fmt"
	"strings"
	"unsafe"
)

// ---------------------------------------------------ContainsFunc 1-----------------------------------

// ---------------------------------------------------ContainsFunc 2-----------------------------------

// ---------------------------------------------------ContainsAny 1-----------------------------------

func ContainsAnyT() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAnyU(s, chars string) bool {
	return strings.IndexAny(s, chars) >= 0
}

// ---------------------------------------------------ContainsAny 2-----------------------------------

// ---------------------------------------------------Contains 2-----------------------------------

func ContainsT() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

func ContainsU(s, substr string) bool {
	return strings.Index(s, substr) >= 0
}

/*
Index(s, substr string) int, 返回 substr 在 s 中第一次出现的位置，不存在返回 -1；采用RabinKarp算法
Index 函数会先对 substr 的长度 n 进行判断，对特殊情况做快速处理。
其次，如果长度 n 以及 len(s) 足够小，则使用BruteForce算法：即暴力匹配，
拿 substr 与 s[i:i+n] 进行比较，如果相等，返回 i，其中 i = (from 0 to len(s) - n)...
最后，会先尝试暴力匹配，如果匹配失败次数超过临界点，则换成 RabinKarp 算法。

func IndexU(s, substr string) int {
	// 特殊的情况处理：
	n := len(substr)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return strings.IndexByte(s, substr[0]) //
	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	case n <= bytealg.MaxLen:
		// Use brute force when s and substr both are small
		if len(s) <= bytealg.MaxBruteForce {
			return bytealg.IndexString(s, substr)
		}
		c0 := substr[0]
		c1 := substr[1]
		i := 0
		t := len(s) - n + 1
		fails := 0
		for i < t {
			if s[i] != c0 {
				// IndexByte is faster than bytealg.IndexString, so use it as long as
				// we're not getting lots of false positives.
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
					return -1
				}
				i += o + 1
			}
			if s[i+1] == c1 && s[i:i+n] == substr {
				return i
			}
			fails++
			i++
			// Switch to bytealg.IndexString when IndexByte produces too many false positives.
			if fails > bytealg.Cutover(i) {
				r := bytealg.IndexString(s[i:], substr)
				if r >= 0 {
					return r + i
				}
				return -1
			}
		}
		return -1
	}

	c0 := substr[0]
	c1 := substr[1]
	i := 0
	t := len(s) - n + 1
	fails := 0
	for i < t {
		if s[i] != c0 {
			o := strings.IndexByte(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if s[i+1] == c1 && s[i:i+n] == substr {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			// See comment in ../bytes/bytes.go.
			// 如果失败次数过多，则使用 indexRabinKarp(s[i:], substr
			j := bytealg.IndexRabinKarp(s[i:], substr)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
}
*/

// ---------------------------------------------------Contains 2-----------------------------------

// UserStringCompareExample ---------------------------------------------------compare 1-----------------------------------
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
