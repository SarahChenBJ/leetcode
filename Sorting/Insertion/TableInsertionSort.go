package insertion

import (
	"github.com/sarahchen/leetcode/dataStruct"
)

const (
	MaxINT = int(^uint(0) >> 1)
)

//LinkSertionSort ..
func LinkInsertionSort(head *dataStruct.ListNode) *dataStruct.ListNode {
	if head == nil {
		return nil
	}

	MX := &dataStruct.ListNode{Val: MaxINT, Next: head}
	p := head.Next
	head.Next = MX
	for p != nil {
		r := MX
		for p.Val >= r.Next.Val {
			r = r.Next
		}
		n, q := r.Next, p.Next
		r.Next = p
		p.Next = n
		p = q
	}
	head = MX.Next
	p = head
	for ; p.Next != MX; p = p.Next {
	}
	p.Next = nil
	return head
}
