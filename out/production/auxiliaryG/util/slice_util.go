package util

// DeleteSliceIndex 删除 Slice 某个下标的号元素
func DeleteSliceIndex[T int | string | bool](a []T, index int) []T {
	copy(a[index:], a[index+1:])
	a = a[:len(a)-1]
	return a
}

// DeleteSliceIndex2 删除 Slice 某个下标的号元素
func DeleteSliceIndex2[T int | string | bool](a []T, index int) []T {
	res := make([]T, len(a)-1)
	copy(res, a[0:index])
	copy(res[index:], a[index+1:])
	return res
}
