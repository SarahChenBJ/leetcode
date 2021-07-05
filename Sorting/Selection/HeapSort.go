package selection

import "fmt"

func HeapSorting(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		HeapAdjust(array, i, len(array)-1)
	}
	for i := len(array) - 1; i >= 0; i-- {
		t := array[i]
		array[i] = array[0]
		array[0] = t
		fmt.Println(i, array)
		HeapAdjust(array, 0, i-1)
	}
}

//HeapAdjust assumes array[s+1:n] follows the heap sorting structure
func HeapAdjust(array []int, s, m int) {
	if s >= m {
		return
	}
	val, i := array[s], 0
	for i = 2*s + 1; i <= m; i = i * 2 {
		if i+1 <= m && array[i] < array[i+1] {
			i++
		}
		if val < array[i] {
			array[s] = array[i]
			s = i
		}
	}
	array[s] = val
}
