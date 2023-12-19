package linkedlist

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := ListNode{
		Val:  9,
		Next: &ListNode{9, &ListNode{9, nil}},
	}
	l2 := ListNode{
		Val:  9,
		Next: &ListNode{9, nil},
	}

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(l3)

	l4 := addTwoNumbers1(&l1,&l2
	fmt.Println(l4)
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

var carry = 0

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	v1, v2 := 0, 0
	l1Next, l2Next := &ListNode{}, &ListNode{}
	if l1 == nil {
		v1 = 0
		l1Next = nil
	} else {
		v1 = l1.Val
		l1Next = l1.Next
	}

	if l2 == nil {
		v2 = 0
		l2Next = nil
	} else {
		v2 = l2.Val
		l2Next = l2.Next
	}

	val := (v1 + v2 + carry) % 10
	carry = (v1 + v2 + carry) / 10
	next := addTwoNumbers(l1Next, l2Next)
	return &ListNode{Val: val, Next: next}
}

// addTwoNumers return a list
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
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
