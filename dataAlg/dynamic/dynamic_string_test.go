package dynamic

import (
	"fmt"
	"testing"
)

func Test_isMatch(t *testing.T) {
	fmt.Println(isMatch("aaa", "bbb"))
}

func isMatch(s string, p string) bool {
	// 1. ini special s+1,p+1
	table := make([][]bool, len(s)+1)
	for i := 0; i < len(table); i++ {
		table[i] = make([]bool, len(p)+1)
	}
	table[0][0] = true

	// 2. int
	for j := 2; j < len(p)+1; j++ {
		if p[j-1] == '*' {
			table[0][j] = table[0][j-2]
		}
	}

	// 3.
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				// 4.
				table[i][j] = table[i-1][j-1]
			} else if p[j-1] == '*' {
				// 5.
				empty := table[i][j-2]
				nonempty := (s[i-1] == p[j-2] || p[j-2] == '.') && table[i-1][j]
				table[i][j] = empty || nonempty
			}
		}
	}

	// 6.
	return table[len(s)][len(p)]
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
