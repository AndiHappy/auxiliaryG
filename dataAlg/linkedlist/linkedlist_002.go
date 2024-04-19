package linkedlist

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (receiver *ListNode) String() string {
	rS := ""
	if receiver != nil {
		rS = rS + strconv.Itoa(receiver.Val)
	}
	if receiver.Next != nil {
		rS = rS + "---> " + receiver.Next.String()
	}
	return rS
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	tmp := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmp.Next = l2
			break
		}
		if l2 == nil {
			tmp.Next = l2
		}
		tmpV := l1.Val + l2.Val + carry
		carry = tmpV / 10
		tmpNode := &ListNode{Val: tmpV % 10}
		tmp.Next = tmpNode
		tmp = tmpNode
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
	}
	if carry != 0 {
		tmp.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	fmt.Println(l1)
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}}
	fmt.Println(l2)
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)
}
