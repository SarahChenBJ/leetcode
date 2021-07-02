package swap

/*
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0075.Sort-Colors/
题目 #
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library’s sort function for this problem.

Example 1:


Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0’s, 1’s, and 2’s, then overwrite array with total number of 0’s, then 1’s and followed by 2’s.
Could you come up with a one-pass algorithm using only constant space?

*/

/*SortColors : 要求一次遍历全部分好, 用快排
[0:i] 0
[i+1:h] 1
[h+1:j] 2
每次出现 0: i和 h 都要挪动
每次出现 1: 只 h 挪动
*/
func SortColors(array []int) {
	i, h := -1, -1
	for j := 0; j < len(array); j++ {
		if array[j] == 0 {
			array[j] = 2
			h++
			array[h] = 1
			i++
			array[i] = 0

		}
		if array[j] == 1 {
			array[j] = 2
			h++
			array[h] = 1

		}
	}
}
