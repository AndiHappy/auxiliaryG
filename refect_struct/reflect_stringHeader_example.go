package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s = "string header"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
	hdata := hdr.Data
	fmt.Println(hdata)
	var ss string = "testValue"
	hdr.Data = uintptr(unsafe.Pointer(&ss)) // case 6 (this case)
	hdr.Len = len(ss)
	fmt.Println(s)

	bytes := []byte("abcdefgh")
	v := unsafe.String(&bytes[0], 3)
	fmt.Println(v)

	bPoint := unsafe.StringData(v)
	v2 := unsafe.String(bPoint, 1)
	fmt.Println(v2)
}
