package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func merge(arr []int, p int, q int, r int) []int {

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

func mergesort(arr []int, p int, r int) []int {
	if r-p == 1 {
		return arr
	}
	if p < r {
		var q = (p + r) / 2
		arr = mergesort(arr, p, q)
		arr = mergesort(arr, q, r)
		arr = merge(arr, p, q, r)

	}

	return arr
}

func main() {
	/*

	   to run:

	   $ go run mergesort.go 23

	*/

	// seed
	rand.Seed(time.Now().Unix())

	// get array size
	arraySize, err := strconv.Atoi(os.Args[1])
	fmt.Println(arraySize)
	if err == nil {

		// generate a random array of integers
		fmt.Println("Initial Array")
		var unsorted = rand.Perm(arraySize)
		fmt.Println(unsorted)

		var sorted = mergesort(unsorted, 0, len(unsorted))
		fmt.Println("sorted array")
		fmt.Println(sorted)

	}
}
