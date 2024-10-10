package util

import "strings"

// If 一种另类的三元计算符
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// ReplaceName 企业名称，除了第一个字符外其余的字符全部替换为*
func ReplaceName(name string) string {
	if len(name) == 0 {
		return ""
	}
	if len(name) == 1 {
		return name
	}
	corpName := []rune(name)
	if len(corpName) > 0 {
		corpName = append(corpName[0:1], []rune(strings.Repeat("*", len(corpName)-1))...)
	}
	return string(corpName)
}
