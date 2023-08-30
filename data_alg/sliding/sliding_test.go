package sliding

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func TestTwoSumE(t *testing.T) {
	fmt.Println(twoSumE([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumE([]int{2, 7, 11, 15}, 9, 90, 900))
}

func twoSumE(nums []int, target int, args ...interface{}) []int {
	if args == nil {
		fmt.Println("true")
		testv(args)
	} else {
		fmt.Printf("%+v \n", args)
		testv(args)
	}

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

func testv(args []interface{}) {
	if args == nil {
		fmt.Println("testv,true")
	} else {
		fmt.Printf("testv,%+v \n", args)
	}
}

// lengthOfLongestSubstring find length of the longest sub string without repeat character
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	max := 0
	from := 0
	si := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		sic := s[i]
		if index, ok := si[sic]; ok {
			if from <= index {
				from = index + 1
			}
		}
		if i-from+1 > max {
			max = i - from + 1
		}
		si[sic] = i
	}
	return max
}
