package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintSlice[T Element](s []T) {
	fmt.Printf("type: %s, kind: %s, len: %v, cap:%v,add:%p,value: %v\n",
		reflect.TypeOf(s), reflect.TypeOf(s).Kind(), len(s), cap(s), &s[0], s)
}

func PrintPara(p any) {
	fmt.Printf("type is %s, kind is %s, len is %v, cap is %v,value is %v, point add:%p \n",
		reflect.TypeOf(p), reflect.TypeOf(p).Kind(), p, &p)
}

// PrintSliceMember 打印 Slice 的成员变量
func PrintSliceMember[T Element](slice []T) {
	if reflect.TypeOf(slice).Kind() == reflect.Slice {
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
		fmt.Printf("addr:%v len:%v cap:%v value:%v \n", hdr.Data, hdr.Len, hdr.Cap, slice)
	}
}
