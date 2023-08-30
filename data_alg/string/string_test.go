package string

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	var p []byte = make([]byte, 1024)
	base64.StdEncoding.Decode(p, []byte("aHR0cHM6Ly90YWdzczAxLnBybyMvcmVnaXN0ZXI/aW52aXRlPVUwNkMwcWtj"))
	fmt.Print(string(p))

	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aa"))
	fmt.Println(longestPalindrome("aaa"))
	fmt.Println(longestPalindrome("aaaa"))

}

func Test_convert(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert_use_strings_Builder("PAYPALISHIRING", 3))
	fmt.Println("PAHNAPLSIIGYIR")

	fmt.Println("-----------------------")

	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert_use_strings_Builder("PAYPALISHIRING", 4))

	fmt.Println("PINALSIGYAHRPI")

	fmt.Println(convert("A", 1))
	fmt.Println(convert_use_strings_Builder("A", 1))
	fmt.Println("A")

}

/*
*
PAYPALISHIRING
P   A   H   N
A P L S I I G
Y   I   R
*/
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows <= 1 {
		return s
	}

	//loop := numRows+numRows-2
	var sBuiler = make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		sBuiler[i] = make([]string, 0)
	}

	for j := 0; j < len(s); {
		for i := 0; i < numRows && j < len(s); i++ {
			sBuiler[i] = append(sBuiler[i], string(s[j]))
			j++
		}

		for i := numRows - 2; i > 0 && j < len(s); i-- {
			sBuiler[i] = append(sBuiler[i], string(s[j]))
			j++
		}
	}

	var ss string
	for i := 0; i < len(sBuiler); i++ {
		for j := 0; j < len(sBuiler[i]); j++ {
			ss = ss + sBuiler[i][j]
		}
	}
	return ss
}

func convert_use_strings_Builder(s string, numRows int) string {
	rows := make([]strings.Builder, numRows+1)

	idx := 0
	for idx < len(s) {
		for i := 0; i < numRows && idx < len(s); i++ {
			rows[i].WriteByte(s[idx])
			idx++
		}
		for j := numRows - 2; 0 < j && idx < len(s); j-- {
			rows[j].WriteByte(s[idx])
			idx++
		}
	}

	var zigzag strings.Builder
	for _, r := range rows {
		zigzag.WriteString(r.String())
	}

	return zigzag.String()
}

// 判断是否是回文字符串
func longestPalindrome(s string) string {
	max := 1
	from, to := 0, 0
	for i := 1; i < len(s); i++ {
		f1, t1, m := judge(i, i, s)
		if m > max {
			from, to = f1, t1
			max = m
		}
		f1, t1, m = judge(i-1, i, s)
		if m > max {
			from, to = f1, t1
			max = m
		}
	}
	return s[from : to+1]
}

func judge(from int, to int, s string) (f, t, m int) {
	for from >= 0 && to < len(s) && s[from] == s[to] {
		from--
		to++
	}
	return from + 1, to - 1, (to - 1) - (from + 1) + 1
}
