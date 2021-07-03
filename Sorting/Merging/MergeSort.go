package merging

import (
	"fmt"

	"github.com/sarahchen/leetcode/sorting/insertion"
)

/*MergingSort is the classic Merging one array*/
func MergingSort(array []int) {
	mergingSortCoreRecursive(array, 0, len(array)-1)
}

/*This is a recursive implementation*/
func mergingSortCoreRecursive(array []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergingSortCoreRecursive(array, start, mid)
	mergingSortCoreRecursive(array, mid+1, end)
	merge(array, start, mid, end)
}

/*MerginigSort_norecursive : This is non-recursive implementation*/
func MerginigSort_norecursive(array []int) {
	for i := 2; i <= 2*(len(array)); i *= 2 {
		//j只能从[0,len(array)-1]
		//edge case 中if j = len(array)-1, mid 会越界(从 i=4 开始, mid 就超出界限了)
		for j := 0; j < len(array)-1; j += i {
			start := j
			end := j + i - 1
			//mid 必须放在 end 判断之前, 不然单数个长度的时候, 最后一次的 mid 取得位置就错了
			mid := (start + end) / 2
			if end >= len(array) {
				end = len(array) - 1
			}
			merge(array, start, mid, end)
		}
		fmt.Println(array, i)
	}
}

func merge(array []int, start, mid, end int) {
	la, ra := []int{}, []int{}
	for i := start; i <= mid; i++ {
		la = append(la, array[i])
	}
	for i := mid + 1; i <= end; i++ {
		ra = append(ra, array[i])
	}
	la = append(la, insertion.MaxINT)
	ra = append(ra, insertion.MaxINT)
	for i, j, k := 0, 0, start; k <= end; {
		if la[i] <= ra[j] {
			array[k] = la[i]
			i++
			k++
		} else {
			array[k] = ra[j]
			j++
			k++
		}
	}
}

/*

https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0088.Merge-Sorted-Array/

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
Constraints:

-10^9 <= nums1[i], nums2[i] <= 10^9
nums1.length == m + n
nums2.length == n
*/

func MergeArrays(a, b []int) {
	lenA, j := findTrueLen(a), 0
	for i, k := 0, 0; k < lenA; {
		for i < lenA && a[i] < b[j] {
			i++
			k++
		}
		s1, s2 := k, j
		for j < len(b) && a[i] >= b[j] {
			k++
			j++
		}
		if s1 != k || s2 != j {
			move := k - s1
			for t := len(a) - 1; t >= k; t-- {
				a[t] = a[t-move]
			}
			for t1, t2 := s1, s2; t1 < k && t2 < j; {
				a[t1] = b[t2]
				t1++
				t2++
			}
			lenA += move
		}
	}
	for i := lenA; i < len(a) && j < len(b); {
		a[i] = b[j]
		i++
		j++
	}
}

func findTrueLen(array []int) int {
	i := len(array) - 1
	for i >= 0 && array[i] == 0 {
		i--
	}
	return i + 1
}
