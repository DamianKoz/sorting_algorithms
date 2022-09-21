package main

import (
	"fmt"
	"testing"
)

func Test_arrays(t *testing.T) {

	t.Run("array initialisation", func(t *testing.T) {
		got := len(initializeArray(5))
		want := 5
		if got != want {
			t.Errorf("Got %v, Want %v", got, want)
		}
	})

	t.Run("array shuffling", func(t *testing.T) {
		arr := initializeArray(5)
		want := shuffle(arr)
		fmt.Printf("\nARRAY %v,\n SHUFFLED: %v\n", arr, want)
		if equal(arr, want) {
			t.Errorf("Got %v, Want %v", arr, want)
		}
	})

}

func equal(arr1, arr2 []int) bool {
	// Return True, if both arrays are deeply equal, with same length and same elements at same index
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
