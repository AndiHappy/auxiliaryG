package auxiliaryGo

import (
	"testing"
)

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func TestStruct(t *testing.T) {
	p := LinkedList{
		Val: 1,
		Next: &LinkedList{
			Val: 2,
			Next: &LinkedList{
				Val: 3,
				Next: &LinkedList{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
}
