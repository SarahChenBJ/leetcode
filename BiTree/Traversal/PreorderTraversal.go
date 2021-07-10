package bitree

import (
	bitree "github.com/sarahchen/leetcode/BiTree"
	"github.com/sarahchen/leetcode/dataStruct"
)

/*https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0144.Binary-Tree-Preorder-Traversal/*/

func Preorder(node *dataStruct.TreeNode) []int {
	ts := new(bitree.TreeStack)
	var array []int
	ts.Push(node)
	for !ts.IsEmpty() {
		p := ts.Pop()
		array = append(array, p.Val)
		if p.Right != nil {
			ts.Push(p.Right)
		}
		if p.Left != nil {
			ts.Push(p.Left)
		}
	}

	return array
}
