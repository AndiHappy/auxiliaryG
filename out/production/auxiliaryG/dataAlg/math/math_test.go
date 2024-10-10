package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	fmt.Println(reverse(123), 321)
	fmt.Println(reverse(-123), -321)
	fmt.Println(reverse(-120), -21)
	fmt.Println(reverse(120), 21)
	fmt.Println(reverse(1534236469), 964634351)
	fmt.Println(reverse(1563847412))
}

func reverse(x int) int {
	flag := false

	var c uint32
	if x < 0 {
		flag = true
		c = uint32(-x)
	} else {
		c = uint32(x)
	}

	var r uint32 = 0
	var tmp uint32 = 0
	for c > 0 {
		tmp = 10*r + c%10
		if (tmp-c%10)/10 != r {
			return 0
		}
		c = c / 10
		r = tmp
	}

	if flag {
		return -int(r)
	} else {
		return int(r)
	}

}
