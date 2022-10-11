package main

import (
	"algorithms-and-datastructures-in-go/utils"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

type DynamicStack struct {
	stackPointer  int
	stackCapacity int
	stack         []*node
}

func newDynamicStack(capacity int) *DynamicStack {
	return &DynamicStack{
		0,
		capacity,
		make([]*node, capacity),
	}
}

func (s *DynamicStack) push(val *node) {
	if s.stackPointer == s.stackCapacity {
		//double stack capacity
		s.stackCapacity *= 2

		var old_stack = s.stack
		s.stack = make([]*node, s.stackCapacity)
		copy(s.stack, old_stack)
	}

	//push
	s.stack[s.stackPointer] = val
	s.stackPointer += 1
}

func (s *DynamicStack) pop() *node {

	if s.stackPointer < s.stackCapacity/3 {
		//deallocate last third
		s.stackCapacity = 2 * s.stackCapacity / 3

		var old_stack = s.stack
		s.stack = make([]*node, s.stackCapacity)

		copy(s.stack, old_stack)
	}

	if s.stackPointer == 0 {
		return nil
	} else {
		s.stackPointer -= 1
		return s.stack[s.stackPointer]
	}
}

////

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

/////

type node struct {
	key   int
	value int
	left  *node
	right *node
}

func makeNode(key int, value int) *node {
	return &node{key: key, value: value}
}

type tree struct {
	root *node
}

func makeEmptyTree() *tree {
	return &tree{}
}

func (t *tree) add(key int, value int) {
	node := makeNode(key, value)

	if t.root == nil {
		t.root = node
	} else {
		parent := t.root
		for {
			if parent.key == node.key { //overwrite
				parent.value = value
				break

			} else if parent.key > node.key { //go left
				if parent.left == nil {
					parent.left = node
					break
				} else {
					parent = parent.left
				}

			} else { //go right
				if parent.right == nil {
					parent.right = node
					break
				} else {
					parent = parent.right
				}
			}
		}
	}
}

func (t *tree) find(key int) *node {

	if t.root == nil {
		return nil
	} else {
		parent := t.root
		for {
			if parent.key == key { //found!
				return parent

			} else if parent.key > key { //go left
				if parent.left == nil {
					return nil
				} else {
					parent = parent.left
				}

			} else { //go right
				if parent.right == nil {
					return nil
				} else {
					parent = parent.right
				}
			}
		}
	}
}

func timeFind(n int, tries int) []time.Duration {

	//create and initialize array with keys for the tries from 0 - biggest in array
	array := generateSorted(n)

	var keys = make([]int, tries)
	for index := range keys {
		keys[index] = rand.Intn(array[len(array)-1] + 1)
	}

	t := makeEmptyTree()

	for _, val := range shuffle(array) {
		t.add(val, val)
	}

	var duration_data []time.Duration

	for _, key := range keys {
		var time_start = time.Now()
		t.find(key)
		var duration = time.Now().Sub(time_start)

		duration_data = append(duration_data, duration)
	}

	return duration_data
}

func (t *tree) dfs() (nextClosure func() *node) {
	stack := newDynamicStack(1000)
	stack.push(t.root)

	return func() *node {
		current := stack.pop()

		if current != nil {
			if current.right != nil {
				stack.push(current.right)
			}

			if current.left != nil {
				stack.push(current.left)
			}
		}

		return current

	}
}

func (t *tree) inOrderTraversal() (nextClosure func() *node) {
	stack := newDynamicStack(1000)
	current := t.root

	gotoLeftmost := func() {
		for (current != nil) && (current.left != nil) {
			stack.push(current)
			current = current.left
		}
	}

	gotoLeftmost()

	return func() *node {
		oldCurrent := current

		if current == nil {
			return nil
		}

		if current.right != nil {
			current = current.right
			gotoLeftmost()
		} else {
			current = stack.pop()
		}

		return oldCurrent
	}
}

func main() {
	treeResult := utils.GetResults(timeFind, 200, utils.IntLinspace(30, 0, 500000))
	arrayResult := utils.GetResults(timeSearchBinary, 200, utils.IntLinspace(30, 0, 500000))

	plot := plot.New()
	plot.Title.Text = "array binary search VS binary tree search \naverage, 200 tries"

	treeMed := treeResult.Scatter(utils.Avg, draw.PlusGlyph{}, utils.Red)
	arrayMed := arrayResult.Scatter(utils.Avg, draw.PlusGlyph{}, utils.Blue)

	plot.Legend.Add("tree find", treeMed)
	plot.Legend.Add("binary search", arrayMed)

	plot.X.Label.Text = "n"
	plot.Y.Label.Text = "runtime (ns)"

	plot.Add(arrayMed, treeMed)
	plot.Save(300, 300, "treeAverage.svg")

}
