package hashmap

import (
	"fmt"
	"testing"
)

// twoSum hashmap time space
// ##leetcode001
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{}
	}
	i2v := make(map[int]int, len(nums))
	for i, v := range nums {
		aTarget := target - v
		if index, ok := i2v[aTarget]; ok {
			return []int{index, i}
		}
		i2v[v] = i
	}
	return []int{}
}

// brute force
func twoSum_bruteforce(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{}
	}

	for i, v := range nums {
		//此次的循环是一个可以优化的内容
		for j, v2 := range nums {
			if v+v2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TestHashMap001(t *testing.T) {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6}, 7))
}
