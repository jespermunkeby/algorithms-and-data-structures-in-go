package main

type linkedList struct {
	root *node
}

type node struct {
	value    int
	next     *node
	previous *node
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

func (list *linkedList) append(other linkedList) { //
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

func (list *linkedList) pop() int { //
	if list.isEmpty() {
		panic("popped from empty stack!")
	}

	val := list.root.value
	list.root = list.root.next

	return val
}

func main() {

}
