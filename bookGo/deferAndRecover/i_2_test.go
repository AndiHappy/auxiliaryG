package main

import (
	"fmt"
	"testing"
)

func TestI2(*testing.T) {
	var whatever [3]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
