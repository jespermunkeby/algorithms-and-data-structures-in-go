package main

import (
	"math/rand"
	"testing"
)

func isSorted(data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if !(data[i] <= data[i+1]) {
			return false
		}
	}

	return true
}

func TestSortSelection(t *testing.T) {
	// some random tests
	for i := 0; i < 100; i++ {
		if !isSorted(sortSelection(rand.Perm(1000))) {
			t.Error("not sorted")
		}
	}

	//some manual tests
	data := [][]int{
		{1, 2, 3, 32, 3, 2, 1, 4, 4, 5},
		{9, 8, 7, 6, 5, 4, 34, 3, 2, 1},
		{6345, 54, 45635, 47, 342547, 2346, 435, 2347, 62345, 6, 247},
		{1, 1, 1, 1, 2, 2, 2, 2, 3, 2, 3, 2, 3, 1, 2, 2, 1, 2, 43, 4, 3, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 2, 3, 2, 34, 4, 4, 4, 4},
	}

	for _, element := range data {
		if !isSorted(sortSelection(element)) {
			t.Error("not sorted")
		}
	}

}

func TestSortInsertion(t *testing.T) {
	// some random tests
	for i := 0; i < 100; i++ {
		if !isSorted(sortInsertion(rand.Perm(100))) {
			t.Error("not sorted")
		}
	}

	//some manual tests
	data := [][]int{
		{1, 2, 3, 32, 3, 2, 1, 4, 4, 5},
		{9, 8, 7, 6, 5, 4, 34, 3, 2, 1},
		{6345, 54, 45635, 47, 342547, 2346, 435, 2347, 62345, 6, 247},
		{1, 1, 1, 1, 2, 2, 2, 2, 3, 2, 3, 2, 3, 1, 2, 2, 1, 2, 43, 4, 3, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 2, 3, 2, 34, 4, 4, 4, 4},
	}

	for _, element := range data {
		if !isSorted(sortInsertion(element)) {
			t.Error("not sorted")
		}
	}

}

func TestSortMerge(t *testing.T) {
	// some random tests
	for i := 0; i < 100; i++ {
		if !isSorted(sortMerge(rand.Perm(1000))) {
			t.Error("not sorted")
		}
	}

	//some manual tests
	data := [][]int{
		{1, 2, 3, 32, 3, 2, 1, 4, 4, 5},
		{9, 8, 7, 6, 5, 4, 34, 3, 2, 1},
		{6345, 54, 45635, 47, 342547, 2346, 435, 2347, 62345, 6, 247},
		{1, 1, 1, 1, 2, 2, 2, 2, 3, 2, 3, 2, 3, 1, 2, 2, 1, 2, 43, 4, 3, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 2, 3, 2, 34, 4, 4, 4, 4},
	}

	for _, element := range data {
		if !isSorted(sortMerge(element)) {
			t.Error("not sorted")
		}
	}

}
