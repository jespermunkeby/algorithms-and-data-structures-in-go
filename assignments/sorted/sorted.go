package main

import (
	"fmt"
	"math/rand"
)

func generateSortedNoDuplicates(n int) []int {
	var arr = make([]int, n)
	var val = 0
	for index := range arr {
		val += rand.Intn(10) + 1
		arr[index] = val
	}

	return arr
}

func shuffle(arr []int) []int {
	perm := rand.Perm(len(arr))
	shuffled := make([]int, len(perm))

	for index, corresponding_index := range perm {
		shuffled[index] = arr[corresponding_index]
	}

	return shuffled
}

func search(arr []int, key int) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}

	return false
}

func searchSorted(arr []int, key int) bool {
	for _, value := range arr {
		if value > key {
			return false
		}
		if value == key {
			return true
		}
	}

	return false
}

func searchBinary(arr []int, key int) bool {
	if len(arr) == 0 { //base case
		return false
	} else if len(arr) == 1 { //base case
		return key == arr[0]
	} else {
		half := (len(arr) - 1) / 2
		elem_half := arr[half]

		if elem_half == key {
			return true
		} else if elem_half > key {
			return searchBinary(arr[:half], key)
		} else {
			return searchBinary(arr[half+1:], key)
		}
	}
}

func main() {

	fmt.Println(searchBinary([]int{1, 3}, 2))

}
