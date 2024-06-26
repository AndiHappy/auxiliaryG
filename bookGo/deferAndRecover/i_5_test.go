package main

import (
	"fmt"
	"testing"
)

func f() (r int) {
	t := 1
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f3() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return 1
}

func TestI5(*testing.T) {
	fmt.Println(f())
	fmt.Println(f2())
	fmt.Println(f3())
}
