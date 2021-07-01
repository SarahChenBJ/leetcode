package insertion

var (
	delta = [3]int{3, 2, 1}
)

//ShellSort ..
func ShellSort(array []int) {
	if len(array) == 0 {
		return
	}
	for _, k := range delta {
		for j := k; j < len(array); j++ {
			val := array[j]
			p := j - k
			for ; p >= 0 && val < array[p]; p -= k {
				array[p+k] = array[p]
			}
			array[p+k] = val
		}
	}
}
