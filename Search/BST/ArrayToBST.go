package BST

import "github.com/sarahchen/leetcode/dataStruct"

/*
108. Convert Sorted Array to Binary Search Tree #
https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0108.Convert-Sorted-Array-to-Binary-Search-Tree/
题目 #
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
*/

func BalancedBSTFromAscArray(array []int, start, end int) *dataStruct.TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	return &dataStruct.TreeNode{
		Val:   array[mid],
		Left:  BalancedBSTFromAscArray(array, start, mid-1),
		Right: BalancedBSTFromAscArray(array, mid+1, end),
	}
}
