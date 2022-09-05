package utils

import "math"

func Median(data []int) float64 {

	var length = len(data)
	var midpoint = float64(length) / 2

	var floored = int(math.Floor(midpoint))
	var ceiled = int(math.Ceil(midpoint))

	if floored == ceiled {
		return float64(data[floored])
	} else {
		return float64(data[floored]+data[floored]) / 2
	}
}

func Average(data []int) float64 {
	var sum = 0
	for _, elem := range data {
		sum += elem
	}

	return float64(sum) / float64(len(data))
}

func Outliers(data []int) (minimum float64, maximum float64) {
	if len(data) == 0 {
		panic("expected data to contain elements")
	}

	var min = data[0]
	var max = data[0]

	for _, elem := range data {
		if elem > max {
			max = elem
		}

		if elem < min {
			min = elem
		}
	}

	return float64(min), float64(max)
}
