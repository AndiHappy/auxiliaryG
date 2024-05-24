package linkedlist002

import (
	"fmt"
	u "github.com/AndiHappy/auxiliaryG/dataAlg/util"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := u.ListNode{
		Val:  9,
		Next: &u.ListNode{Val: 9, Next: &u.ListNode{9, nil}},
	}
	l2 := u.ListNode{
		Val:  9,
		Next: &u.ListNode{Val: 9},
	}

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(l3)

	l4 := addTwoNumbers1(&l1, &l2)
	fmt.Println(l4)
}

/**
 * Definition for singly-linked list.
 */

var carry = 0

// addTwoNumbers
// leetcode002
func addTwoNumbers(l1 *u.ListNode, l2 *u.ListNode) *u.ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	val := (JudgeValue(l1) + JudgeValue(l2) + carry) % 10
	carry = (JudgeValue(l1) + JudgeValue(l2) + carry) / 10
	next := addTwoNumbers(NextNode(l1), NextNode(l2))
	return &u.ListNode{Val: val, Next: next}
}

// addTwoNumbers3
// leetcode002
func addTwoNumbers3(l1 *u.ListNode, l2 *u.ListNode) *u.ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	val := (JudgeValue(l1) + JudgeValue(l2) + carry) % 10
	carry = (JudgeValue(l1) + JudgeValue(l2) + carry) / 10
	next := addTwoNumbers(NextNode(l1), NextNode(l2))
	return &u.ListNode{Val: val, Next: next}
}

func NextNode(l *u.ListNode) *u.ListNode {
	if l == nil {
		return nil
	} else {
		return l.Next
	}
}

func JudgeValue(l *u.ListNode) int {
	if l == nil {
		return 0
	} else {
		return l.Val
	}
}

// addTwoNumers return a list
func addTwoNumbers1(l1 *u.ListNode, l2 *u.ListNode) *u.ListNode {
	dump := &u.ListNode{}
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
		next := u.ListNode{Val: v % 10}
		dump.Next = &next
		dump = dump.Next
	}
	if carry > 0 {
		dump.Next = &u.ListNode{
			carry, nil,
		}
	}
	return result.Next
}
