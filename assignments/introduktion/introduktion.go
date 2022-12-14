package main

import (
	"algorithms-and-datastructures-in-go/utils"
	"fmt"
	"math/rand"
	"time"
)

func timeOfArrayAssignment(n int, tries int) []time.Duration {
	var array = make([]int, n)
	var random_array_index = rand.Intn(n)

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		var time_start = time.Now()
		array[random_array_index] = random_array_index
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeOfArrayAccess(n int, tries int) []time.Duration {
	var array = make([]int, n)
	var random_array_index = rand.Intn(n)

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		var time_start = time.Now()
		_ = array[random_array_index]
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func arraySearch(array []int, key int) bool {
	for _, elem := range array {
		if elem == key {
			return true
		}
	}

	return false
}

func timeOfArraySearch(n int, tries int) []time.Duration {
	//create and initialize array with random keys from 0 - n*10
	var array = make([]int, n)
	for index := range array {
		array[index] = rand.Intn(n * 10)
	}

	//create and initialize array with keys for the tries from 0 - n*10
	var keys = make([]int, tries)
	for index := range keys {
		keys[index] = rand.Intn(n * 10)
	}

	var duration_data []time.Duration

	for _, key := range keys {
		var time_start = time.Now()
		arraySearch(array, key)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func ArrayDuplicateSearch(a []int, b []int) bool {
	for _, elem_a := range a {
		for _, elem_b := range b {
			if elem_a == elem_b {
				return true
			}
		}
	}

	return false
}

func timeOfArrayDuplicateSearch(n int, tries int) []time.Duration {
	var a, b = make([]int, n), make([]int, n)
	for index := range a {
		a[index] = rand.Intn(n * 10)
	}
	for index := range b {
		b[index] = rand.Intn(n * 10)
	}

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		var time_start = time.Now()
		ArrayDuplicateSearch(a, b)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data

}

func computeAndDisplay(timetester_func utils.TimeTester, tries int, n_samples []int) {
	for _, n := range n_samples {
		var nanoseconds []int
		for _, elem := range timetester_func(n, tries) {
			nanoseconds = append(nanoseconds, int(elem.Nanoseconds()))
		}

		//display
		var avg = utils.Average(nanoseconds)
		var median = utils.Median(nanoseconds)
		var minimum, maximum = utils.Outliers(nanoseconds)

		fmt.Println(
			"min: ", minimum, "\t",
			"max: ", maximum, "\t",
			"avg: ", avg, "\t",
			"med: ", median, "\t",
			"n: ", n,
		)
	}

}

func main() {

	fmt.Println("array access")
	computeAndDisplay(
		timeOfArrayAccess,
		20000,
		[]int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000},
	)
	fmt.Println("---")

	fmt.Println("array write")
	computeAndDisplay(
		timeOfArrayAssignment,
		20000,
		[]int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000},
	)
	fmt.Println("---")

	// fmt.Println("array search")
	// computeAndDisplay(
	// 	timeOfArraySearch,
	// 	20000,
	// 	[]int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000},
	// )
	// fmt.Println("---\n")

	fmt.Println("array write")
	computeAndDisplay(
		timeOfArrayDuplicateSearch,
		10,
		[]int{1000000, 2000000, 3000000, 4000000, 5000000, 6000000, 7000000, 8000000},
	)
	fmt.Println("---")
}
