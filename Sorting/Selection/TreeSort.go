package selection

//CampaignSorting ..
func CampaignSorting(array []int) {
	for i := 0; i < len(array); i++ {

		for j, k := i, len(array)-1; j < k; {
			p := j
			for ; p <= (j+k)/2; p++ {
				if array[p] > array[k-p+j] {
					t := array[p]
					array[p] = array[k-p+j]
					array[k-p+j] = t
				}

			}
			k = p - 1
		}

	}
}
