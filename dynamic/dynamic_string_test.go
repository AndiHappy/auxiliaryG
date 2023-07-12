package dynamic

import (
	"fmt"
	"testing"
)

func Test_isMatch(t *testing.T) {
	fmt.Println(isMatch("aaa", "bbb"))
}

// isMatch 自动匹配
func isMatch(s, p string) bool {
	return false
}

func Test_isMatch_simple(t *testing.T) {
	fmt.Println(isMatch_simple("aaa", ".a."))
}

// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like . or *.
// first assign p only contains only lowercase letters a-z, and characters like .
// s [0,i] 匹配 p [0,j]
func isMatch_simple(s, p string) bool {
	if p == "" {
		return (s == "" || len(s) == 0)
	} else {
		return isMatch_simple_dynamic(s, p, 0, 0)
	}
}

func isMatch_simple_dynamic(s, p string, ss, ps int) bool {
	dynamicA := make([][]bool, len(s))
	for i := 0; i < len(p); i++ {
		if (i < len(s) && i < len(p) && s[i] == p[i]) || p[i] == '.' || (i > 0 && dynamicA[0][i-1] && p[i] == '.') {
			dynamicA[0][i] = true
		}
	}

	fmt.Println(dynamicA)

	return dynamicA[len(s)-1][len(p)-1]
}
