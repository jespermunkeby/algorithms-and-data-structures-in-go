package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSearchBinary(t *testing.T) {
	// some random tests for TRUE
	for n := 1; n < 100; n++ {
		arr := generateSorted(n)
		key := arr[rand.Intn(n)]
		if !searchBinary(arr, key) {
			fmt.Print(arr, key, "\n")
			t.Errorf("false negative")
		}
	}

	//some particular tests for FALSE
	for key, arr := range [][]int{
		{1, 2, 3, 4},          //0
		{0, 2, 3, 4},          //1
		{0, 1, 3, 4},          //2
		{0, 1, 2, 4, 4, 5, 6}, //3
	} {
		if searchBinary(arr, key) {
			fmt.Print(arr, key, "\n")
			t.Errorf("false positive")
		}
	}
}
