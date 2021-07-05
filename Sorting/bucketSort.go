package sorting

import swap "github.com/sarahchen/leetcode/Sorting/Swap"

const BucketSize = 10

func BucketSort(array []int) {
	bucketSortCore(array, BucketSize)
}

func bucketSortCore(array []int, bucketSize int) {
	if len(array) == 0 {
		return
	}

	//get the max and min
	min, max := array[0], array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}

	//set up the buckets
	bucketCount := (max-min)/bucketSize + 1
	buckets := make([][]int, bucketCount)

	//fill the buckets
	for i := 0; i < len(array); i++ {
		buckets[(array[i]-min)/bucketSize] = append(buckets[(array[i]-min)/bucketSize], array[i])
	}

	//sorting every bucket and writing it back to array
	for i, j := 0, 0; i < bucketCount; i++ {
		barr := buckets[i]
		swap.QuickSort(barr)
		for k := 0; k < len(barr); k++ {
			array[j] = barr[k]
			j++
		}
	}
}
