package BiTree

import (
	"fmt"

	"github.com/sarahchen/leetcode/dataStruct"
)

/*https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0144.Binary-Tree-Preorder-Traversal/*/

type TreeStack struct {
	Stack []*dataStruct.TreeNode
}

func Preorder(node *dataStruct.TreeNode) []int {
	ts := new(TreeStack)
	var array []int
	p := node
	for p != nil || !ts.IsEmpty() {
		fmt.Println(p)
		for p != nil {
			ts.Push(p)
			p = p.Left
		}
		q := ts.Pop()
		array = append(array, q.Val)
		if q.Right != nil {
			p = q.Right
		}
	}
	return array
}

func (t *TreeStack) Push(n *dataStruct.TreeNode) {
	if n != nil {
		t.Stack = append(t.Stack, n)
	}
}
func (t *TreeStack) Pop() *dataStruct.TreeNode {
	if len(t.Stack) > 0 {
		n := t.Stack[len(t.Stack)-1]
		t.Stack = t.Stack[:len(t.Stack)-1]
		return n
	}
	return nil
}
func (t *TreeStack) IsEmpty() bool {
	return len(t.Stack) == 0
}
