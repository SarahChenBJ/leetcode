package bitree

import (
	"fmt"

	bitree "github.com/sarahchen/leetcode/BiTree"
	"github.com/sarahchen/leetcode/dataStruct"
)

/*https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0094.Binary-Tree-Inorder-Traversal/*/

func Inorder(node *dataStruct.TreeNode) []int {
	ts := new(bitree.TreeStack)
	
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
