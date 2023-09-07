package gosdkexample

import (
	"fmt"
	"unsafe"
)

// StringDataT returns a pointer to the underlying bytes of str.
// For an empty string the return value is unspecified, and may be nil.
// Since Go strings are immutable, the bytes returned by StringData
// must not be modified.
func StringDataT() {
	s := "hello,kkoman"
	sBytes := unsafe.StringData(s)
	fmt.Println(*sBytes)
}

func StringDataU() {
	// 直接通过强转string(bytes)或者[]byte(str)会带来数据的复制，性能不佳，所以在追求极致性能场景，
	//我们会采用『骇客』的方式，来实现这两种类型的转换,比如k8s采用下面的方式：
	fmt.Println(K8sCloneString{}.toBytes("ddd"))
}
