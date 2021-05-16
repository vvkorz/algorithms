package sorting

// QuickSortPartition is a routine that splits an array into 2 subarrays
// by returning an index of an element where each element below this index
// is less than element at this index
func QuickSortPartition(arr []int, p, r int) int {
	var pivot int = arr[r]
	i := p - 1 // place of the pivot element in rearranged array
	for j := p; j < r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // swap elements putting the smaller to the left of pivot
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1] // put pivot in its place
	return i + 1
}

// QuickSort is an implememntation of a quicksort algorithm
func QuickSort(arr []int, p, r int) []int {
	if p < r {
		q := QuickSortPartition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
	return arr
}
