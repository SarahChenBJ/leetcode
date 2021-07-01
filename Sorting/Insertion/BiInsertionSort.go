package insertion

//BiInsertionSort ..
func BiInsertionSort(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		k := array[i+1]
		low, high := 0, i
		for low <= high {
			mid := (low + high) / 2
			if k == array[mid] {
				low = mid + 1
				break
			}
			if k > array[mid] {
				low = mid + 1
				continue
			}
			if k < array[mid] {
				high = mid - 1
				continue
			}
		}
		for j := i; j >= low; j-- {
			array[j+1] = array[j]
		}
		array[low] = k
	}
}
