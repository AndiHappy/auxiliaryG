package main

import (
	"fmt"
	"reflect"
)

func TypeAndValueAccess() {
	var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("value:", reflect.ValueOf(x).Interface())

	type MyInt int
	var xx MyInt = 7
	v := reflect.ValueOf(xx)
	fmt.Println(v.Kind())           //int类型
	fmt.Println(reflect.TypeOf(xx)) //main.MyInt
}

func StructReflectSet() {
	t := T{23, "skidoo"}
	//新的方法 Elem returns a type's element type.
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)

	t1 := T{23, "skidoo"}
	//如果不是&t是什么情况？
	s1 := reflect.ValueOf(t1).Elem() //panic: reflect: call of reflect.Value.Elem on struct Value
	fmt.Println(s1)
}

type T struct {
	A int
	B string
}

// 1. Reflection goes from interface value to reflection object.
func main2() {
	v := reflect.ValueOf(new(float32))
	v = v.Elem()
	switch v.Kind() {
	case reflect.Float32:
		v.SetFloat(4.56)
	}
	s := v.Float()
	fmt.Println(s)

	v1 := reflect.ValueOf(new(string))
	v1 = v1.Elem()
	switch v1.Kind() {
	case reflect.String:
		v1.SetString("4.56")
	}
	s1 := v1.String()
	fmt.Println(s1)
}
