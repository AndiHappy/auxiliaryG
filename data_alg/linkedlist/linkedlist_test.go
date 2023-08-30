package linkedlist

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := ListNode{
		Val:  2,
		Next: &ListNode{3, &ListNode{4, nil}},
	}
	l2 := ListNode{
		Val:  4,
		Next: &ListNode{5, &ListNode{6, nil}},
	}

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(l3)
}

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	if p.Next != nil {
		return strconv.Itoa(p.Val) + " -> " + p.Next.String()
	} else {
		return strconv.Itoa(p.Val)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := &ListNode{}
	result := dump
	carry := 0
	for l1 != nil || l2 != nil {
		l1v := 0
		if l1 != nil {
			l1v = l1.Val
			l1 = l1.Next
		}
		l2v := 0
		if l2 != nil {
			l2v = l2.Val
			l2 = l2.Next
		}
		v := carry + l1v + l2v
		carry = v / 10
		next := ListNode{Val: v % 10}
		dump.Next = &next
		dump = dump.Next
	}
	if carry > 0 {
		dump.Next = &ListNode{
			carry, nil,
		}
	}
	return result.Next
}
