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

func twoSum(nums []int, target int, args ...interface{}) []int {
	fmt.Printf("dd:%+v", args)
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
