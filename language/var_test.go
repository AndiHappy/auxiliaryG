package gotutorial

import (
	"fmt"
	"testing"
)

func TestVarExample(t *testing.T) {
	var a *int
	*a = 100 // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100 //panic: assignment to entry in nil map
	fmt.Println(b)
}
