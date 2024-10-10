package auxiliaryGo

import "strconv"

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
