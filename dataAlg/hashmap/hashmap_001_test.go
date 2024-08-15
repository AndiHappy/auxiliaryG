package hashmap

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

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
func twoSumBruteForce(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{}
	}
	for i, v := range nums {
		//内循环是一个可以优化的内容，声明map就是为了能够取消该内循环
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
