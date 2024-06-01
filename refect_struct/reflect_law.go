package main

import (
	"fmt"
	"reflect"
)

// 1. Reflection goes from interface value to reflection object.
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("value:", reflect.ValueOf(x).Interface())

	type MyInt int
	var xx MyInt = 7
	v := reflect.ValueOf(xx)
	fmt.Println(v.Kind())           //int类型
	fmt.Println(reflect.TypeOf(xx)) //main.MyInt
}
