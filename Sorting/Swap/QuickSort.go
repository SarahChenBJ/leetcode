package swap

/*
1. 选最后一个元素为 pix
2. 通过一趟交换, 维持 i,j:
(1) array[0:i] < pix
(2) array[i+1:j] > pix
3. 一次遍历结束之后, 把 pix <--> array[i+1]
这样, array 就被 pix 分为两部分 [0:i-1] [i+1,len]
*/

func partition(array []int, start, end int) int {
	if start < 0 || end >= len(array) {
		return -1
	}
	pix, i := array[end], start-1
	for j := start; j < end; j++ {
		if array[j] < pix {
			i++
			a := array[i]
			array[i] = array[j]
			array[j] = a
		}
	}
	i++
	a := array[i]
	array[i] = array[end]
	array[end] = a
	return i
}

func quickSortCore(array []int, start, end int) {
	if start >= end {
		return
	}
	p := partition(array, start, end)
	quickSortCore(array, start, p-1)
	quickSortCore(array, p+1, end)
}

//QuickSort ..
func QuickSort(array []int) {
	quickSortCore(array, 0, len(array)-1)
}
