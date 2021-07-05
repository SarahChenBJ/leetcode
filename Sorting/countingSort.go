package sorting

//CountingSort :
//pay more attention to : buckets[array[i]-min]++
func CountingSort(array []int) {
	if len(array) == 0 {
		return
	}
	//find the max and min to determin the bucket size
	min, max := array[0], array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}
	bucketLen := max - min + 1
	buckets := make([]int, bucketLen)
	for i := 0; i < len(array); i++ {
		buckets[array[i]-min]++
	}

	//fill buckets
	for i, j := 0, 0; i < len(buckets); i++ {
		for count := buckets[i]; count > 0; count-- {
			array[j] = i + min
			j++
		}
	}

}
