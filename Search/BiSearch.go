package search

/* 简单二分查找*/
func BiSearch(array []int, k int) int {
	ps := -1
	for low, high := 0, len(array)-1; low <= high; {
		mid := (low + high) / 2
		if k == array[mid] {
			ps = mid
			break
		}
		if k > array[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ps
}

/*查找第一个与 target 相等的元素，时间复杂度 O(logn)*/
func BiSearchTheFirst(array []int, k int) int {
	ps := -1
	for low, high := 0, len(array)-1; low <= high; {
		mid := (low + high) / 2

		if k > array[mid] {
			low = mid + 1
		}
		if k < array[mid] {
			high = mid - 1
		}
		if k == array[mid] {
			if mid == 0 || array[mid-1] < k {
				ps = mid
				break
			}
			high = mid - 1
		}
	}
	return ps
}

/*查找最后一个与 target 相等的元素，时间复杂度 O(logn)*/
func BiSearchTheLast(array []int, k int) int {
	ps := -1
	for low, high := 0, len(array)-1; low <= high; {
		mid := (low + high) / 2
		if k > array[mid] {
			low = mid + 1
		}
		if k < array[mid] {
			high = mid - 1
		}
		if k == array[mid] {
			if mid == len(array)-1 || array[mid+1] > k {
				ps = mid
				break
			}
			low = mid + 1
		}
	}
	return ps
}

/*查找第一个大于等于 target 的元素，时间复杂度 O(logn)*/
func BiSearchTheFirstLarger(array []int, k int) int {
	ps := -1
	for low, high := 0, len(array)-1; low <= high; {
		mid := (low + high) / 2
		if array[mid] > k {
			if mid == 0 || array[mid-1] <= k {
				return mid
			}
			high = mid - 1
		}
		if array[mid] == k {
			if mid == 0 || array[mid+1] > k {
				return mid
			}
			low = mid + 1
		}
		if array[mid] < k {
			low = mid + 1
		}
	}
	return ps
}

/*查找最后一个小于等于 target 的元素，时间复杂度 O(logn)*/
