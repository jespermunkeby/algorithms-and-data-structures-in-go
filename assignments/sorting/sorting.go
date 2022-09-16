package main

import (
	"algorithms-and-datastructures-in-go/utils"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

func sortSelection(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		candindate := i

		for j := i; j < len(data); j++ {
			if data[j] < data[candindate] {
				candindate = j
			}
		}

		//swap
		temp := data[i]
		data[i] = data[candindate]
		data[candindate] = temp
	}

	return data
}

func sortInsertion(data []int) []int {
	for i := 1; i < len(data); i++ {

		for swap_cursor := i; (swap_cursor >= 1) && (data[swap_cursor] < data[swap_cursor-1]); swap_cursor-- {
			//swap
			temp := data[swap_cursor]
			data[swap_cursor] = data[swap_cursor-1]
			data[swap_cursor-1] = temp
		}
	}
	return data
}

func merge(a []int, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i := 0
	for (len(a) > 0) && (len(b) > 0) {
		if a[0] < b[0] {
			merged[i] = a[0]
			a = a[1:]
		} else {
			merged[i] = b[0]
			b = b[1:]
		}

		i++
	}

	for len(a) > 0 {
		merged[i] = a[0]
		a = a[1:]

		i++
	}

	for len(b) > 0 {
		merged[i] = b[0]
		b = b[1:]

		i++
	}
	return merged
}

func sortMerge(data []int) []int {
	if len(data) == 1 {
		return data
	} else {
		return merge(
			sortMerge(data[:len(data)/2]),
			sortMerge(data[len(data)/2:]),
		)
	}
}

func timeSortSelection(n int, tries int) []time.Duration {
	var durationData []time.Duration

	for i := 0; i < tries; i++ {
		data := rand.Perm(n)

		var time_start = time.Now()
		sortSelection(data)
		var duration = time.Now().Sub(time_start)

		durationData = append(durationData, duration)
	}

	return durationData
}

func timeSortInsertion(n int, tries int) []time.Duration {
	var durationData []time.Duration

	for i := 0; i < tries; i++ {
		data := rand.Perm(n)

		var time_start = time.Now()
		sortInsertion(data)
		var duration = time.Now().Sub(time_start)

		durationData = append(durationData, duration)
	}

	return durationData
}

func timeSortMerge(n int, tries int) []time.Duration {
	var durationData []time.Duration

	for i := 0; i < tries; i++ {
		data := rand.Perm(n)

		var time_start = time.Now()
		sortMerge(data)
		var duration = time.Now().Sub(time_start)

		durationData = append(durationData, duration)
	}

	return durationData
}

func main() {
	exp := int(math.Pow(10, 3))
	insertionResults := utils.GetResults(timeSortInsertion, 10, []int{
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
	})

	selectionResults := utils.GetResults(timeSortSelection, 10, []int{
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
	})

	duplicatePlot := plot.New()
	duplicatePlot.Title.Text = "median runtime, 10 tries"

	insertionMedian := insertionResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)
	selectionMedian := selectionResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Blue)

	duplicatePlot.Add(insertionMedian, selectionMedian)
	duplicatePlot.Legend.Add("insertion sort", insertionMedian)
	duplicatePlot.Legend.Add("selection sort", selectionMedian)
	duplicatePlot.X.Label.Text = "n (array length)"
	duplicatePlot.Y.Label.Text = "runtime (ns)"

	duplicatePlot.Save(500, 500, "insertion_and_selection.svg")

	///////////

	exp = int(math.Pow(10, 5))

	results := utils.GetResults(timeSortMerge, 10, []int{
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
	})

	plot := plot.New()
	plot.Title.Text = "runtime of merge sort, 10 tries"
	median := results.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)
	plot.Add(median)
	plot.Legend.Add("median", median)
	plot.X.Label.Text = "n (array length)"
	plot.Y.Label.Text = "runtime (ns)"

	plot.Save(500, 500, "merge.svg")

}
