package main

import (
	"fmt"
	"github.com/AndiHappy/helpGo/util"
)

// slice数组的变化
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]
	util.PrintSliceMember(s2)
	s2 = append(s2, 100)
	util.PrintSliceMember(s2)
	s1[2] = 20
	fmt.Println(s1)
	fmt.Println(s2)
	s2 = append(s2, 200)
	util.PrintSliceMember(s2)
	s1[2] = -20
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
