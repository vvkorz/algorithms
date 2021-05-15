package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/vvkorz/algorithms/sorting"
)

func main() {
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

		var sorted = sorting.MergeSort(unsorted, 0, len(unsorted))
		fmt.Println("MergeSort array")
		fmt.Println(sorted)

	}
}
