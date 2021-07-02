package swap

//BubbleSort ..
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				a := array[j]
				array[j] = array[j+1]
				array[j+1] = a
			}
		}
	}
}
