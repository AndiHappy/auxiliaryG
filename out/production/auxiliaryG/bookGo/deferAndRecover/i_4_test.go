package main

import (
	"fmt"
	"testing"
)

func TestI4(*testing.T) {
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
