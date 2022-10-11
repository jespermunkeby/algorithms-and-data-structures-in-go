package main

func wraparound(n int, bound int) int {
	if n < 0 {
		return bound - ((-n) % bound)
	} else {
		return n % bound
	}
}

type queue struct {
	arr     []int
	dynamic bool
	start   int
	length  int
}

func makeDynamicQueue() *queue {
	return &queue{
		arr:     make([]int, 1),
		dynamic: true,
		start:   0,
		length:  0,
	}
}
func makeStaticQueue(capacity int) *queue {
	return &queue{
		arr:     make([]int, capacity),
		dynamic: false,
		start:   0,
		length:  0,
	}
}

func (q *queue) isEmpty() bool { return q.length == 0 }

func (q *queue) isFull() bool { return q.length == len(q.arr) }

func (q *queue) add(val int) {
	if q.isFull() {
		if q.dynamic { //reallocate
			oldArr := q.arr
			q.arr = make([]int, len(q.arr)*2)
			for i := range oldArr {
				q.arr[i] = wraparound(oldArr[i+q.start], len(oldArr))
			}
		} else {
			panic("added to full static queue")
		}
	}

	q.arr[wraparound(q.start+q.length, len(q.arr))] = val
	q.length += 1
}

func (q *queue) remove() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}

	val := q.arr[q.start]

	q.start = wraparound(q.start+1, len(q.arr))
	q.length--

	return val, true
}

func (q *queue) display() {
	for i, e := range q.arr {
		if i == q.start {
			print("[S] ")
		}
		if i == wraparound(q.start+q.length, len(q.arr)) {
			print("[E] ")
		}
		print(e, " ")

	}

	print("\n")
}

func main() {
	q := makeStaticQueue(10)
	q.display()

	q.add(0)
	q.add(1)
	q.add(2)
	q.add(3)
	q.add(4)
	q.add(5)

	q.display()

	q.remove()
	q.remove()
	q.remove()

	q.display()

	q.add(17)
	q.add(18)
	q.add(19)

	q.display()

	q.add(5)
	q.add(5)
	q.add(5)
	q.add(5)

	q.display()
}
