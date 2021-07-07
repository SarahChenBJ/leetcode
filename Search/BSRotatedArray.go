package search

import "fmt"

/*
33. Search in Rotated Sorted Array #
题目 #
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm’s runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/

func SearchInRotatedArray(array []int, target int) int {
	index, pivot := -1, -1
	low, high := 0, len(array)-1
	for low <= high {
		mid := (low + high) / 2
		if mid == len(array)-1 {
			pivot = mid
			break
		}
		if array[mid] <= array[mid+1] {
			low = mid + 1
		}
		if array[mid] > array[mid+1] {
			if mid >= 1 && array[mid] > array[mid-1] {
				pivot = mid
				break
			}
			high = mid
		}
	}
	if pivot == -1 {
		low, high = 0, len(array)-1
	} else {
		if target == array[pivot] {
			low, high = pivot, pivot
		} else if target < array[0] {
			low, high = pivot+1, len(array)-1
		} else {
			low, high = 0, pivot
		}
	}

	fmt.Println(pivot, low, high)
	for low <= high {
		mid := (low + high) / 2
		if target == array[mid] {
			index = mid
			break
		} else if target > array[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return index
}
