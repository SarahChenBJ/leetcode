package bitree

import (
	"github.com/sarahchen/leetcode/dataStruct"
)

type TreeStack struct {
	Stack []*dataStruct.TreeNode
	flag  []bool
}

func (t *TreeStack) Push(n *dataStruct.TreeNode) {
	if n != nil {
		t.Stack = append(t.Stack, n)
		t.flag = append(t.flag, false)
	}
}

func (t *TreeStack) Pop() *dataStruct.TreeNode {
	if len(t.Stack) > 0 {
		n := t.Stack[len(t.Stack)-1]
		nlen := len(t.Stack) - 1
		t.Stack = t.Stack[:nlen]
		t.flag = t.flag[:nlen]
		return n
	}
	return nil
}

func (t *TreeStack) IsEmpty() bool {
	return len(t.Stack) == 0
}

func (t *TreeStack) GetTop() (*dataStruct.TreeNode, bool) {
	if len(t.Stack) > 0 {
		return t.Stack[len(t.Stack)-1], t.flag[len(t.Stack)-1]
	}
	return nil, false
}

func (t *TreeStack) MarkedTop(flag bool) {
	if len(t.Stack) > 0 {
		t.flag[len(t.Stack)-1] = flag
	}
}

type Queue struct{ queue []*dataStruct.TreeNode }

func (q *Queue) PopHead() *dataStruct.TreeNode {
	l := len(q.queue)
	if l > 0 {
		n := q.queue[0]
		if l >= 1 {
			q.queue = q.queue[1:l]
		}
		return n
	}
	return nil
}

func (q *Queue) PushTail(n *dataStruct.TreeNode) {
	if n != nil {
		q.queue = append(q.queue, n)
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
