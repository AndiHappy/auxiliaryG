package main

import (
	"fmt"
	"github.com/AndiHappy/auxiliaryG/refect_struct/custom_struct"
	"reflect"
	"runtime"
	"unsafe"
)

func read_foo(f *custom_struct.Foo) {
	v := reflect.ValueOf(*f)
	y := v.FieldByName("y")
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println(y.Interface()) // panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
}

func change_foo(f *custom_struct.Foo) {
	// Since structs are organized in memory order, we can advance the pointer
	// by field size until we're at the desired member. For y, we advance by 8
	// since it's the size of an int on a 64-bit machine and the int "x" is first
	// in the representation of Foo.
	// If you wanted to alter x, you wouldn't advance the pointer at all, and simply
	// would need to convert ptrTof to the type (*int)
	ptrTof := unsafe.Pointer(f)
	ptrTof = unsafe.Pointer(uintptr(ptrTof) + uintptr(8)) // Or 4, if this is 32-bit
	ptrToy := (**custom_struct.Foo)(ptrTof)
	*ptrToy = nil // or *ptrToy = &Foo{} or whatever you want
}

func main() {
	f := custom_struct.NewFoo(1, 2)
	//read_foo(&f)
	change_foo(&f)
	fmt.Println(f)
}

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}
