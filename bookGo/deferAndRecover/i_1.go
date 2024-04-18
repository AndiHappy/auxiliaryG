package main

import (
	"fmt"
)

var arrP = []int{1}

func main() {
	myfunc1(arrP)     // 0,2
	fmt.Println(arrP) //1
	arrP = append(arrP, 3)
	arrP = append(arrP, 4)
	myfunc2(arrP)     //[9 3 4 5]
	fmt.Println(arrP) //[9 3 4]
}

func myfunc1(arr []int) {
	arr = append(arr, 2)
	arr[0] = 0
	fmt.Println(arr)
	return
}

func myfunc2(arr []int) {
	arr = append(arr, 5)
	arr[0] = 9
	fmt.Println(arr)
	return
}
