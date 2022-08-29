package main

import (
	"fmt"

	"math/rand"

	algs "github.com/DamianKoz/sorting_algorithms/algorithms"
)

func main() {
	// arr := []int{5, 4, 2, 6, 7, 4}
	arr := initializeArray(1000000, 1000000)
	fmt.Printf("Unsorted Array: \n %v \n", arr)
	sortedArray := algs.BubbleSort(arr)
	fmt.Printf("Sorted Array \n %v", sortedArray)
}

func initializeArray(maxNum, length int) []int {
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(maxNum))
	}
	return arr
}
