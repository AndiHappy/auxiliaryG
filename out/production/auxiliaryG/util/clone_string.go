package util

import "unsafe"

// K8sCloneString 主要使用的是unsafe的Pointer的方法
type K8sCloneString struct {
}

// toBytes performs unholy acts to avoid allocations
func (K8sCloneString) toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// toString performs unholy acts to avoid allocations
func (K8sCloneString) toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (K8sCloneString) SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func (K8sCloneString) StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
