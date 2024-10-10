package sliding_window

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abbbbbbcdedga"))
}

// lengthOfLongestSubstring find length of the longest
// sub string without repeat character
// leetcode003
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var max = 0
	var from = 0
	//循环遍历string的时候，s[index]是一个uint8
	si := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		sic := s[i]
		//如果碰到重复了，需要首先判断是否滑动窗口
		if preI, ok := si[sic]; ok {
			if from <= preI {
				from = preI + 1
			} else { //如果是from小于preI是什么情况下，发生的！
				fmt.Println(s)
			}
		}
		if i-from+1 > max {
			max = i - from + 1
		}
		si[sic] = i
	}
	return max
}
