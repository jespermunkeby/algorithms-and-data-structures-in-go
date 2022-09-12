package main

import (
	"algorithms-and-datastructures-in-go/utils"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

func generateSorted(n int) []int {
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

func timeSearch(n int, tries int) []time.Duration {
	//create and initialize array with random keys from 0 - n*10
	var array = shuffle(generateSorted(n))

	//create and initialize array with keys for the tries from 0 - biggest in array
	var keys = make([]int, tries)
	for index := range keys {
		keys[index] = rand.Intn(array[len(array)-1] + 1)
	}

	var duration_data []time.Duration

	for _, key := range keys {
		var time_start = time.Now()
		search(array, key)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeSearchSorted(n int, tries int) []time.Duration {
	//create and initialize array with random keys from 0 - n*10
	var array = generateSorted(n)

	//create and initialize array with keys for the tries from 0 - biggest in array
	var keys = make([]int, tries)
	for index := range keys {
		keys[index] = rand.Intn(array[len(array)-1] + 1)
	}

	var duration_data []time.Duration

	for _, key := range keys {
		var time_start = time.Now()
		searchSorted(array, key)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeSearchBinary(n int, tries int) []time.Duration {
	//create and initialize array with random keys from 0 - n*10
	var array = generateSorted(n)

	//create and initialize array with keys for the tries from 0 - biggest in array
	var keys = make([]int, tries)
	for index := range keys {
		keys[index] = rand.Intn(array[len(array)-1] + 1)
	}

	var duration_data []time.Duration

	for _, key := range keys {
		var time_start = time.Now()
		searchBinary(array, key)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func searchDuplicate(a []int, b []int) bool {
	for _, elem_a := range a {
		for _, elem_b := range b {
			if elem_a == elem_b {
				return true
			}
		}
	}

	return false
}

func timeSearchDuplicate(n int, tries int) []time.Duration {

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		a := shuffle(generateSorted(n))
		b := shuffle(generateSorted(n))

		var time_start = time.Now()
		searchDuplicate(a, b)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data

}

func searchBinaryDuplicate(a []int, b []int) bool {
	for _, elem_a := range a {
		if searchBinary(b, elem_a) {
			return true
		}
	}

	return false
}

func timeSearchBinaryDuplicate(n int, tries int) []time.Duration {

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {

		a := shuffle(generateSorted(n))
		b := generateSorted(n)

		var time_start = time.Now()
		searchBinaryDuplicate(a, b)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data

}

func finalDuplicateSearch(a []int, b []int) bool {
	for _, elem_a := range a {
		if len(b) == 0 {
			return false
		} else if b[0] == elem_a {
			return true
		} else if b[0] < elem_a {
			b = b[1:]
		}
	}

	return false
}

func timeFinalDuplicate(n int, tries int) []time.Duration {

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {

		a := generateSorted(n)
		b := generateSorted(n)

		var time_start = time.Now()
		finalDuplicateSearch(a, b)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func main() {

	// exp := int(math.Pow(10, 4))

	// unsortedSearchResults := utils.GetResults(timeSearch, 50, []int{
	// 	1 * exp,
	// 	2 * exp,
	// 	3 * exp,
	// 	4 * exp,
	// 	5 * exp,
	// 	6 * exp,
	// 	7 * exp,
	// 	8 * exp,
	// 	9 * exp,
	// 	10 * exp,
	// 	11 * exp,
	// 	12 * exp,
	// 	13 * exp,
	// })

	// sortedVsUnsortedPlot := plot.New()
	// sortedVsUnsortedPlot.Title.Text = "sorted vs. unsorted search, 50 tries"

	// medianUnsorted := unsortedSearchResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)
	// averageUnsorted := unsortedSearchResults.Scatter(utils.Avg, draw.CircleGlyph{}, utils.Red)

	// sortedVsUnsortedPlot.Add(medianUnsorted, averageUnsorted)
	// sortedVsUnsortedPlot.Legend.Add("unsorted median", medianUnsorted)
	// sortedVsUnsortedPlot.Legend.Add("unsorted average", averageUnsorted)

	// sortedSearchResults := utils.GetResults(timeSearchSorted, 50, []int{
	// 	1 * exp,
	// 	2 * exp,
	// 	3 * exp,
	// 	4 * exp,
	// 	5 * exp,
	// 	6 * exp,
	// 	7 * exp,
	// 	8 * exp,
	// 	9 * exp,
	// 	10 * exp,
	// 	11 * exp,
	// 	12 * exp,
	// 	13 * exp,
	// })

	// medianSorted := sortedSearchResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Blue)
	// averageSorted := sortedSearchResults.Scatter(utils.Avg, draw.CircleGlyph{}, utils.Blue)

	// sortedVsUnsortedPlot.Add(medianSorted, averageSorted)
	// sortedVsUnsortedPlot.Legend.Add("sorted median", medianSorted)
	// sortedVsUnsortedPlot.Legend.Add("sorted average", averageSorted)

	// sortedVsUnsortedPlot.X.Label.Text = "n (array length)"
	// sortedVsUnsortedPlot.Y.Label.Text = "runtime (ns)"

	// xFunc := plotter.NewFunction(func(x float64) float64 { return x })
	// sortedVsUnsortedPlot.Add(xFunc)
	// sortedVsUnsortedPlot.Legend.Add("y(x) = x", xFunc)

	// sortedVsUnsortedPlot.Save(500, 500, "sortedVSunsorted.svg")

	// ///////////////////////////////////////////////
	// ///////////////////////////////////////////////
	// ///////////////////////////////////////////////

	// exp = int(math.Pow(10, 6)) / 4

	// searchBinaryResults := utils.GetResults(timeSearchBinary, 50, []int{
	// 	1,
	// 	1 * exp,
	// 	2 * exp,
	// 	3 * exp,
	// 	4 * exp,
	// 	5 * exp,
	// 	6 * exp,
	// 	7 * exp,
	// 	8 * exp,
	// 	9 * exp,
	// 	10 * exp,
	// 	11 * exp,
	// 	12 * exp,
	// 	13 * exp,
	// 	14 * exp,
	// 	15 * exp,
	// 	16 * exp,
	// 	17 * exp,
	// 	18 * exp,
	// 	19 * exp,
	// 	20 * exp,
	// 	21 * exp,
	// 	22 * exp,
	// 	23 * exp,
	// 	24 * exp,
	// 	25 * exp,
	// 	26 * exp,
	// 	27 * exp,
	// 	28 * exp,
	// 	29 * exp,
	// 	30 * exp,
	// 	31 * exp,
	// 	32 * exp,
	// })

	// binaryPlot := plot.New()
	// binaryPlot.Title.Text = "binary search, 50 tries"

	// binaryPlot.X.Label.Text = "n (array length)"
	// binaryPlot.Y.Label.Text = "runtime (ns)"

	// medianBinary := searchBinaryResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)
	// averageBinary := searchBinaryResults.Scatter(utils.Avg, draw.CircleGlyph{}, utils.Red)

	// binaryPlot.Add(medianBinary, averageBinary)
	// binaryPlot.Legend.Add("median", medianBinary)
	// binaryPlot.Legend.Add("average", averageBinary)

	// f := func(x float64) float64 { return 430 * math.Log(x*0.0000013) }
	// println(f(65000000))
	// logxFunc := plotter.NewFunction(f)
	// binaryPlot.Add(logxFunc)
	// binaryPlot.Legend.Add("log(x)", logxFunc)

	// binaryPlot.Save(500, 500, "binary.svg")

	// ///////////////////////////////////////////////
	// ///////////////////////////////////////////////
	// ///////////////////////////////////////////////

	exp := int(math.Pow(10, 1)) * 2

	duplicateSearchResults := utils.GetResults(timeSearchDuplicate, 500, []int{
		1,
		1 * exp,
		2 * exp,
		3 * exp,
		4 * exp,
		5 * exp,
		6 * exp,
		7 * exp,
		8 * exp,
		9 * exp,
		10 * exp,
		11 * exp,
		12 * exp,
		13 * exp,
		14 * exp,
		15 * exp,
		16 * exp,
		17 * exp,
		18 * exp,
		19 * exp,
		20 * exp,
		21 * exp,
		22 * exp,
		23 * exp,
		24 * exp,
		25 * exp,
		26 * exp,
		27 * exp,
		28 * exp,
		29 * exp,
		30 * exp,
		31 * exp,
		32 * exp,
		33 * exp,
		34 * exp,
		35 * exp,
		36 * exp,
		37 * exp,
		38 * exp,
		39 * exp,
		40 * exp,
	})

	duplicatePlot := plot.New()
	duplicatePlot.Title.Text = "median runtime of duplicate search, 500 tries"

	medianOld := duplicateSearchResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)
	//maxOld := duplicateSearchResults.Scatter(utils.Max, draw.PyramidGlyph{}, utils.Red)

	duplicatePlot.Add(medianOld)
	duplicatePlot.Legend.Add("old", medianOld)
	// duplicatePlot.Legend.Add("old max", maxOld)

	duplicateBinarySearchResults := utils.GetResults(timeSearchBinaryDuplicate, 500, []int{
		1,
		1 * exp,
		2 * exp,
		3 * exp,
		4 * exp,
		5 * exp,
		6 * exp,
		7 * exp,
		8 * exp,
		9 * exp,
		10 * exp,
		11 * exp,
		12 * exp,
		13 * exp,
		14 * exp,
		15 * exp,
		16 * exp,
		17 * exp,
		18 * exp,
		19 * exp,
		20 * exp,
		21 * exp,
		22 * exp,
		23 * exp,
		24 * exp,
		25 * exp,
		26 * exp,
		27 * exp,
		28 * exp,
		29 * exp,
		30 * exp,
		31 * exp,
		32 * exp,
		33 * exp,
		34 * exp,
		35 * exp,
		36 * exp,
		37 * exp,
		38 * exp,
		39 * exp,
		40 * exp,
	})

	medianBinaryS := duplicateBinarySearchResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Blue)
	//maxBinary := duplicateBinarySearchResults.Scatter(utils.Max, draw.PyramidGlyph{}, utils.Blue)

	duplicatePlot.Add(medianBinaryS)
	duplicatePlot.Legend.Add("binary", medianBinaryS)
	//duplicatePlot.Legend.Add("binary max", maxBinary)

	finalSearchResults := utils.GetResults(timeFinalDuplicate, 500, []int{
		1,
		1 * exp,
		2 * exp,
		3 * exp,
		4 * exp,
		5 * exp,
		6 * exp,
		7 * exp,
		8 * exp,
		9 * exp,
		10 * exp,
		11 * exp,
		12 * exp,
		13 * exp,
		14 * exp,
		15 * exp,
		16 * exp,
		17 * exp,
		18 * exp,
		19 * exp,
		20 * exp,
		21 * exp,
		22 * exp,
		23 * exp,
		24 * exp,
		25 * exp,
		26 * exp,
		27 * exp,
		28 * exp,
		29 * exp,
		30 * exp,
		31 * exp,
		32 * exp,
		33 * exp,
		34 * exp,
		35 * exp,
		36 * exp,
		37 * exp,
		38 * exp,
		39 * exp,
		40 * exp,
	})

	medianFinal := finalSearchResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Green)
	//maxBinary := duplicateBinarySearchResults.Scatter(utils.Max, draw.PyramidGlyph{}, utils.Blue)

	duplicatePlot.Add(medianFinal)
	duplicatePlot.Legend.Add("final", medianFinal)

	duplicatePlot.X.Label.Text = "n (array length)"
	duplicatePlot.Y.Label.Text = "runtime (ns)"

	duplicatePlot.Save(500, 500, "duplicates.svg")

}
