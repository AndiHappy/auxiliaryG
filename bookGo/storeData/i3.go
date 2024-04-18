package main

import "fmt"

func myAppend(s []int) []int {
	// 这里 s 结构体虽然改变了，但并不会改变外层函数的 s 结构体
	s = append(s, 100)
	return s
}
func myAppendPtr(s *[]int) {
	// 会改变外层 s 结构体本身
	*s = append(*s, 100)
	return
}

// slice作为方法的参数，
func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)
	fmt.Println(s)
	fmt.Println(newS)
	s = newS
	myAppendPtr(&s)
	fmt.Println(s)
}
