package main

import (
	"errors"
	"fmt"
	"testing"
)

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
func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer3 error")
	return
}

func TestI6(*testing.T) {
	e1() //nil
	e2() //defer2 error
	e3() //nil
}
