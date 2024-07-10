package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	//a := Spit("  foo bar  baz  ")
	//fmt.Printf("%q\v", a)
	//b := "rererre"
	//fmt.Println(CopyString(b))

	fmt.Printf("%d", 0x7F)

	makeASCIISet("  foo bar  baz  HHKYIYL，llda，打啊")

}

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
// This allocates a total of 32 bytes even though the upper half
// is unused to avoid bounds checks in asciiSet.contains.
type asciiSet [8]uint32 //8个uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c/32] |= 1 << (c % 32)
	}
	return as, true
}

func CopyString(s string) string {
	// Just return a copy.
	return string(append([]byte(nil), s...))
}

func CopyArrayByte(s []byte) []byte {
	// Just return a copy.
	return append([]byte(nil), s...)
}

func TestByteArray2RuneArray() {
	b := "we中！@……"
	runes := ByteArray2RuneArray([]byte(b))
	runes2 := ByteArray2RuneArray2([]byte(b))
	i := 0
	for {
		if i < len(runes2) && i < len(runes) && runes[i] == runes2[i] {
			fmt.Printf("i: %d value:%c  value2:%c\n", i, runes[i], runes2[i])
		} else {
			break
		}
		i++
	}
}

// ByteArray2RuneArray2
// utf8.RuneSelf == 0x80 128 判断是否需要转化
// 利用utf8.DecodeRune解析一字节
func ByteArray2RuneArray2(s []byte) []rune {
	//不能够直接转化：Cannot convert an expression of the type '[]byte' to the type '[]rune'
	//但是可以间接的转
	runes := make([]rune, 0)
	i := 0
	for i < len(s) {
		size := 1
		r := rune(s[i])
		if r >= utf8.RuneSelf {
			r, size = utf8.DecodeRune(s[i:])
		}
		runes = append(runes, r)
		i += size
	}
	return runes
}

func ByteArray2RuneArray(byteArray []byte) []rune {
	//不能够直接转化：Cannot convert an expression of the type '[]byte' to the type '[]rune'
	//但是可以间接的转
	stringV := string(byteArray)
	r := []rune(stringV)
	return r
}

// asciiSpace，分离ascii，这是一个数组，一个256的数组
var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Spit 主要想记录这个循环遍历和跳过特殊的情况，抽象具体逻辑内容
/**
	//首先是处理开头的情况
	segmentStart := 0
	for i < len(s) && conditionFunc(s) {
		i++ //隔离处不满足条件的开始值
	}
    segmentStart := 0
	//正式处理
	for i < len(s) {
		if !conditionFunc(s) {
            i++
			continue
        }
		//处理segmentStart -- i 这段内容
		dealFunction(segmentStart, i)
		i++
		// 再次使用开头的逻辑，处理需要跳过的逻辑
		for i < len(s) && conditionFunc(s) {
			i++
		}
		segmentStart = i //再次确定下一段的开头
	}
    //最后处理最后一个segment
    if segmentStart < len(s) {
		dealFunction(segmentStart, len(s))
    }
*/
func Spit(asciiOnly string) [][]byte {
	s := []byte(asciiOnly)
	a := make([][]byte, 0)
	fieldStart := 0
	i := 0
	// Skip spaces in the front of the input.
	// 首先跳过空格
	for i < len(s) && asciiSpace[s[i]] != 0 {
		i++
	}
	fieldStart = i
	for i < len(s) {
		// continue skip spaces.
		//如果是非空格的情况下，积累i
		if asciiSpace[s[i]] == 0 {
			i++
			continue
		}
		a = append(a, s[fieldStart:i:i])
		i++
		// Skip spaces in between fields.
		for i < len(s) && asciiSpace[s[i]] != 0 {
			i++
		}
		fieldStart = i
	}
	if fieldStart < len(s) { // Last field might end at EOF.
		a = append(a, s[fieldStart:len(s):len(s)])
	}
	return a
}

// IsSeparator 记录是否是一个隔离符号
func IsSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r < utf8.RuneSelf {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
