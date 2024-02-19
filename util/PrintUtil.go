package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintPara(p any) {
	fmt.Printf("type is %s, kind is %s, point add:%p \n", reflect.TypeOf(p), reflect.TypeOf(p).Kind(), &p)
}

// PrintSliceMember 打印 Slice 的成员变量
func PrintSliceMember[T Element](slice []T) {
	if reflect.TypeOf(slice).Kind() == reflect.Slice {
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
		fmt.Printf("addr:%v len:%v cap:%v value:%v \n", hdr.Data, hdr.Len, hdr.Cap, slice)
	}
}
