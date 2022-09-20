package main

import (
	"algorithms-and-datastructures-in-go/utils"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

type linkedList struct {
	root *node
}

type node struct {
	value int
	next  *node
}

func makeLinkedList(root *node) *linkedList {
	return &linkedList{
		root,
	}
}

func makeEmptyLinkedList() *linkedList {
	return &linkedList{}
}

func makeLinkedListFromArray(array []int) *linkedList {
	list := makeEmptyLinkedList()
	for i, _ := range array {
		list.push(array[len(array)-1-i])
	}

	return list
}

func (list *linkedList) isEmpty() bool {
	return list.root == nil
}

func (list *linkedList) last() *linkedList {
	if list.isEmpty() {
		panic("tried to get last on empty list")
	} else {
		for node := list.root; ; node = node.next {
			if node.next == nil {
				return makeLinkedList(node)
			}
		}
	}
}

func (list *linkedList) append(other linkedList) {
	if list.isEmpty() {
		list.root = other.root
	} else {
		list.last().root.next = other.root
	}
}

func (list *linkedList) prepend(other linkedList) {
	other.append(*list)
	list.root = other.root
}

func (list *linkedList) push(val int) {
	list.prepend(*makeLinkedList(&node{value: val}))
}

func (list *linkedList) pop() int {
	if list.isEmpty() {
		panic("popped from empty stack!")
	}

	val := list.root.value
	list.root = list.root.next

	return val
}

type dynArray struct {
	array []int
}

func makeDynArray(n int) *dynArray {
	return &dynArray{make([]int, n)}
}

func (a *dynArray) append(other dynArray) {
	arr := make([]int, len(a.array)+len(other.array))
	i := 0
	for _, e := range a.array {
		arr[i] = e
		i++
	}

	for _, e := range other.array {
		arr[i] = e
		i++
	}

	*a = dynArray{array: arr}
}

func timeAppendArray(n int, tries int) []time.Duration {
	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		//create constantly sized list
		a := makeDynArray(100000)
		//create dynamicallt sized
		b := makeDynArray(n)

		var time_start = time.Now()
		b.append(*a)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeMakeArray(n int, tries int) []time.Duration {
	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		var time_start = time.Now()
		makeDynArray(n)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeAppendNToConstant(n int, tries int) []time.Duration {

	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		//create constantly sized list
		a := makeLinkedListFromArray(rand.Perm(300))
		//create dynamicallt sized
		b := makeLinkedListFromArray(rand.Perm(n))

		var time_start = time.Now()
		a.append(*b)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeAppendConstantToN(n int, tries int) []time.Duration {
	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		//create constantly sized list
		a := makeLinkedListFromArray(rand.Perm(100000))
		//create dynamicallt sized
		b := makeLinkedListFromArray(rand.Perm(n))

		var time_start = time.Now()
		b.append(*a)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func timeMakeLinked(n int, tries int) []time.Duration {
	var duration_data []time.Duration

	for i := 0; i < tries; i++ {
		arr := rand.Perm(n)
		var time_start = time.Now()
		makeLinkedListFromArray(arr)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func main() {
	appendConstantResults := utils.GetResults(timeAppendConstantToN, 10, utils.IntLinspace(30, 0, 50000))

	appendConstantPlot := plot.New()
	appendConstantPlot.Title.Text = "runtime of appending to list of size n, 10 tries"

	acMed := appendConstantResults.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)

	appendConstantPlot.Add(acMed)
	appendConstantPlot.Legend.Add("median", acMed)

	appendConstantPlot.X.Label.Text = "n"
	appendConstantPlot.Y.Label.Text = "runtime (ns)"

	appendConstantPlot.Save(300, 300, "appendConstant.svg")

	//////

	appendDynamicResults := utils.GetResults(timeAppendNToConstant, 100, utils.IntLinspace(50, 0, 10000))

	appendDynamicPlot := plot.New()
	appendDynamicPlot.Title.Text = "runtime of appending list of size n, 100 tries"

	adAvg := appendDynamicResults.Scatter(utils.Avg, draw.PlusGlyph{}, utils.Red)

	appendDynamicPlot.Add(adAvg)
	appendDynamicPlot.Legend.Add("average", adAvg)

	appendDynamicPlot.X.Label.Text = "n"
	appendDynamicPlot.Y.Label.Text = "runtime (ns)"

	appendDynamicPlot.Save(300, 300, "appendDynamic.svg")

	//// Bench against array ops APPEND

	arrayResultsAppend := utils.GetResults(timeAppendArray, 10, utils.IntLinspace(30, 0, 50000))

	linkedVSArrayAppendPlot := plot.New()
	linkedVSArrayAppendPlot.Title.Text = "linked list VS array append, median, 10 tries.\n (array reallocates every append)"

	arrMed := arrayResultsAppend.Scatter(utils.Med, draw.PlusGlyph{}, utils.Blue)

	linkedVSArrayAppendPlot.Add(arrMed)
	linkedVSArrayAppendPlot.Legend.Add("array", arrMed)
	linkedVSArrayAppendPlot.Add(acMed)
	linkedVSArrayAppendPlot.Legend.Add("linked list", acMed)

	linkedVSArrayAppendPlot.X.Label.Text = "n"
	linkedVSArrayAppendPlot.Y.Label.Text = "runtime (ns)"

	linkedVSArrayAppendPlot.Save(300, 300, "arrayVSlinkedAppend.svg")

	//// Bench against array ops MAKE

	arrayResultsMake := utils.GetResults(timeMakeArray, 10, utils.IntLinspace(30, 0, 50000))
	linkedResultsMake := utils.GetResults(timeMakeLinked, 10, utils.IntLinspace(30, 0, 50000))

	linkedVSArrayMakePlot := plot.New()
	linkedVSArrayMakePlot.Title.Text = "linked list VS array init, median, 10 tries."

	arrMedMake := arrayResultsMake.Scatter(utils.Med, draw.PlusGlyph{}, utils.Blue)
	linkedMedMake := linkedResultsMake.Scatter(utils.Med, draw.PlusGlyph{}, utils.Red)

	linkedVSArrayMakePlot.Add(arrMedMake)
	linkedVSArrayMakePlot.Legend.Add("array", arrMedMake)
	linkedVSArrayMakePlot.Add(linkedMedMake)
	linkedVSArrayMakePlot.Legend.Add("linked list", linkedMedMake)

	linkedVSArrayMakePlot.X.Label.Text = "n"
	linkedVSArrayMakePlot.Y.Label.Text = "runtime (ns)"

	linkedVSArrayMakePlot.Save(300, 300, "arrayVSlinkedMake.svg")
}
