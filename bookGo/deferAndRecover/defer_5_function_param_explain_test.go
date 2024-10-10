package main

import (
	"fmt"
	"testing"
)

func pushAnalytic(a int) {
	fmt.Println(a)
}

// Output:
// 10
// That’s because when you use the defer statement,
// it grabs the values right then.
// This is called “capture by value.”
// So, the value of a that gets sent to pushAnalytic is set to 10
// when the defer is scheduled, even though a changes later.
func TestDeferExeOrder(t *testing.T) {
	a := 10
	defer func(a int) {
		pushAnalytic(a)
	}(a)
	a = 20
}

// modify
// output:20
func TestDeferExeOrder_1(t *testing.T) {
	a := 10
	defer func() {
		pushAnalytic(a)
	}()
	a = 20
}

// modify2
// output:20
func TestDeferExeOrder_2(t *testing.T) {
	a := 10
	defer pushAnalytic1(&a)
	a = 20
}

func pushAnalytic1(a *int) {
	fmt.Println(*a)
}

// test method
type Data struct {
	a int
}

func (d Data) pushAnalytic() {
	fmt.Println(d.a)
}

func TestDeferExeOrder_3(t *testing.T) {
	d := Data{a: 10}
	defer d.pushAnalytic() //equal to: defer Data.pushAnalytic(d)
	d.a = 20
}
