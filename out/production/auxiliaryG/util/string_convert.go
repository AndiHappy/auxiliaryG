package util

import (
	"unicode"
	"unicode/utf8"
)

// S2ByteArray 类型转化
func S2ByteArray(s string) []byte {
	return []byte(s)
}

// IsSeparator 记录是否是一个隔离符号
func IsSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r < utf8.RuneSelf {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
