package selection

//SimpleSelection ..
func SimpleSelection(array []int) {
	for i := 0; i < len(array); i++ {
		n := i
		for j := i + 1; j < len(array); j++ {
			if array[n] > array[j] {
				n = j
			}
		}
		a := array[i]
		array[i] = array[n]
		array[n] = a
	}
}
