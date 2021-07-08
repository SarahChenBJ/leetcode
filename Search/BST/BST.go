package BST

import (
	"github.com/sarahchen/leetcode/dataStruct"
)

//SearchBST search k in the tree
//if k is in the tree:     will return true, the treenode
//if k is not in the tree: will return false, the last treenode for the following insertion
func SearchBST(node, lnode *dataStruct.TreeNode, k int) (bool, *dataStruct.TreeNode) {
	if node == nil {
		return false, lnode
	}
	if k == node.Val {
		return true, node
	}
	if k > node.Val {
		return SearchBST(node.Right, node, k)
	}
	return SearchBST(node.Left, node, k)
}

func InsertBST(array []int) *dataStruct.TreeNode {
	if len(array) == 0 {
		return nil
	}
	tree := &dataStruct.TreeNode{Val: array[0]}

	for _, v := range array {
		if k, n := SearchBST(tree, nil, v); !k && n != nil {
			vn := &dataStruct.TreeNode{Val: v}
			if v > n.Val {
				n.Right = vn
			} else {
				n.Left = vn
			}
		}
	}
	return tree
}
