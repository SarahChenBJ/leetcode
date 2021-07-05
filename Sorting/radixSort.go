package sorting

import (
	"strings"
)

/*RadixSortLSD .. */
func RadixSortLSD(array []string) {
	maxLen := maxLen(array)
	//the buckets[0] will store short string
	buckets := make([][]string, 27)
	for i := maxLen - 1; i >= 0; i-- {
		for j, v := range array {
			index := 0
			if i < len([]rune(strings.ToLower(v))) {
				ichar := []rune(strings.ToLower(v))[i]
				index = int(ichar-'a') % 27
			}
			buckets[index] = append(buckets[index], array[j])
		}
		for k, h := 0, 0; k < len(buckets); k++ {
			if len(buckets[k]) != 0 {
				for l := range buckets[k] {
					array[h] = buckets[k][l]
					h++
				}
			}
			buckets[k] = nil
		}
	}
}

func maxLen(array []string) int {
	ml := 0
	for _, v := range array {
		l := len([]rune(v))
		if l > ml {
			ml = l
		}
	}

	return ml
}

/*
164. Maximum Gap #
https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0164.Maximum-Gap/
题目 #
Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:


Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.

Example 2:


Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.

Note:

You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
Try to solve it in linear time/space.
*/
