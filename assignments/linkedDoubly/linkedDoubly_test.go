package main

import "testing"

func TestLast(t *testing.T) {
	a := makeLinkedList(&node{
		value: 1,
		next:  nil,
	})

	if a.last().root.value != 1 {
		panic("incorrect last()")
	}

	a = makeLinkedList(&node{
		value: 1,
		next: &node{
			value: 2,
			next:  nil,
		},
	})

	if a.last().root.value != 2 {
		panic("incorrect last()")
	}

	a = makeLinkedList(&node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next:  nil,
			},
		},
	})

	if a.last().root.value != 3 {
		panic("incorrect last()")
	}

	if a.last().root.next != nil {
		panic("incorrect last()")
	}

	a.last().root.next = &node{
		value: 22,
		next:  nil,
	}

	if a.last().root.value != 22 {
		panic("incorrect last()")
	}

}

func TestAppend(t *testing.T) {
	a := makeEmptyLinkedList()

	b := makeLinkedList(&node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next:  nil,
			},
		},
	})

	a.append(*b)

	if a.last().root.value != 3 {
		panic("incorrect append")
	}

	a = makeLinkedList(&node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next:  nil,
			},
		},
	})

	b = makeLinkedList(&node{
		value: 42,
		next:  nil,
	})

	a.append(*b)
	println(a.last().root.value)

	if a.last().root.value != 42 {
		panic("incorrect append")
	}

}

func TestPrepend(t *testing.T) {
	a := makeEmptyLinkedList()

	b := makeLinkedList(&node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next:  nil,
			},
		},
	})

	a.prepend(*b)

	if a.last().root != b.last().root {
		panic("incorrect prepend")
	}

	a = makeLinkedList(&node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next:  nil,
			},
		},
	})

	b = makeLinkedList(&node{
		value: 4,
		next: &node{
			value: 5,
			next: &node{
				value: 6,
				next:  nil,
			},
		},
	})

	b.prepend(*a)

	root := b.root
	if !(root.value == 1 &&
		root.next.value == 2 &&
		root.next.next.value == 3 &&
		root.next.next.next.value == 4 &&
		root.next.next.next.next.value == 5 &&
		root.next.next.next.next.next.value == 6) {
		panic("incorrect prepend")
	}

}

func TestPopPush(t *testing.T) {
	l := makeEmptyLinkedList()

	l.push(1)
	l.push(2)
	l.push(3)
	l.push(4)
	l.push(5)
	l.push(6)

	if 6 != l.pop() {
		panic("incorrect pop/push")
	}

	if 5 != l.pop() {
		print(l.pop())
		panic("incorrect pop/push")
	}

	if 4 != l.pop() {
		panic("incorrect pop/push")
	}

	if 3 != l.pop() {
		panic("incorrect pop/push")
	}

	if 2 != l.pop() {
		panic("incorrect pop/push")
	}

	if 1 != l.pop() {
		panic("incorrect pop/push")
	}
}
