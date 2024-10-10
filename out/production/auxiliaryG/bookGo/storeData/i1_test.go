package main

import (
	"fmt"
)

// slice底层数据的变化
func main() {
	data := [...]int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5} //数组
	slice := data[0:1:7]                           // Slice data[low, high, max]
	fmt.Println(slice, len(slice), cap(slice))

	slice = data[1:1:7] // Slice data[low, high, max]
	fmt.Println(slice, len(slice), cap(slice))

	slice = data[2:2:7] // Slice data[low, high, max]
	fmt.Println(slice, len(slice), cap(slice))

	//s2 := data[1:100]//Invalid slice index '100' (out of bounds for the 10-element array)
	//s3 := data[1:5:1] //Invalid index values, must be low <= high <= max

	s2 := data[1:10]
	fmt.Println(s2, len(s2), cap(s2))

	s22 := data[1:7]
	fmt.Println(s22, len(s22), cap(s22))

	s3 := data[1:5:5]
	fmt.Println(s3, len(s3), cap(s3))
}
