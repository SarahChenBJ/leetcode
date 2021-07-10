package bitree

import (
	bitree "github.com/sarahchen/leetcode/BiTree"
	"github.com/sarahchen/leetcode/dataStruct"
)

/*https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0145.Binary-Tree-Postorder-Traversal/*/

func Postorder(node *dataStruct.TreeNode) []int {
	ts := new(bitree.TreeStack)
	var array []int
	ts.Push(node)
	for !ts.IsEmpty() {
		p, ok := ts.GetTop()
		for ok {
			array = append(array, p.Val)
			ts.Pop()
			p, ok = ts.GetTop()
		}
		if p != nil && !ok {
			ts.MarkedTop(true)
			if p.Right != nil {
				ts.Push(p.Right)
			}
			if p.Left != nil {
				ts.Push(p.Left)
			}
		}
	}
	return array
}
