package sorting

// Merge is an implementation of a merging routine in the MergeSort algorithm
func Merge(arr []int, p int, q int, r int) []int {

	const sentinel int = 99999

	left_size := q - p
	right_size := r - q

	// initiate 2 new arrays
	var left = make([]int, left_size+1)
	var right = make([]int, right_size+1)

	// copy arrays
	for i := 0; i < left_size; i++ {
		left[i] = arr[p+i]
	}
	for j := 0; j < right_size; j++ {
		right[j] = arr[q+j]
	}
	// put sentinel at the end of the array
	left[left_size] = sentinel
	right[right_size] = sentinel

	var i int = 0
	var j int = 0
	// start merging
	for k := p; k < r; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}

// MergeSort implements merge sort algorithms for sorting
func MergeSort(arr []int, p int, r int) []int {
	if r-p == 1 {
		return arr
	}
	if p < r {
		var q = (p + r) / 2
		arr = MergeSort(arr, p, q)
		arr = MergeSort(arr, q, r)
		arr = Merge(arr, p, q, r)

	}

	return arr
}
