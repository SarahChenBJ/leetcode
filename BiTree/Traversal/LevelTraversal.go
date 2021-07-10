package bitree

import (
	bitree "github.com/sarahchen/leetcode/BiTree"
	"github.com/sarahchen/leetcode/dataStruct"
)

/*https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0102.Binary-Tree-Level-Order-Traversal/*/

func LevelTraversal(node *dataStruct.TreeNode) []int {
	q := new(bitree.Queue)
	var array []int
	q.PushTail(node)
	for !q.IsEmpty() {
		n := q.PopHead()

		array = append(array, n.Val)
		if n.Left != nil {
			q.PushTail(n.Left)
		}
		if n.Right != nil {
			q.PushTail(n.Right)
		}
	}

	return array
}
