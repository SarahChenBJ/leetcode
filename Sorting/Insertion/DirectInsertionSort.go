package insertion

import (
	"fmt"
)

// DirectInsertionSort ...
func DirectInsertionSort(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		j, k := i, array[i+1]
		for ; j >= 0 && k < array[j]; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = k
	}
	for _, val := range array {
		fmt.Print(val)
	}
}
