package main

import (
	"fmt"
	"time"
)

type itemType int

const (
	Add itemType = iota
	Sub
	Mul
	Div
	Mod
	SpecialMult
	Value
)

type item struct {
	itemType itemType
	value    int
}

type stack interface {
	push(int)
	pop() int
}

type staticStack struct {
	stackPointer  int
	stackCapacity int
	stack         []int
}

func newStaticStack(capacity int) *staticStack {
	return &staticStack{
		0,
		capacity,
		make([]int, capacity),
	}
}

func (s *staticStack) push(val int) {
	if s.stackPointer == s.stackCapacity {
		panic("pushed to full stack")
	} else {
		s.stack[s.stackPointer] = val
		s.stackPointer += 1
	}
}

func (s *staticStack) pop() int {
	if s.stackPointer == 0 {
		panic("popped from empty stack")
	} else {
		s.stackPointer -= 1
		return s.stack[s.stackPointer]
	}
}

type DynamicStack struct {
	stackPointer  int
	stackCapacity int
	stack         []int
}

func newDynamicStack(capacity int) *DynamicStack {
	return &DynamicStack{
		0,
		capacity,
		make([]int, capacity),
	}
}

func (s *DynamicStack) push(val int) {
	if s.stackPointer == s.stackCapacity {
		//double stack capacity
		s.stackCapacity *= 2

		var old_stack = s.stack
		s.stack = make([]int, s.stackCapacity)
		copy(s.stack, old_stack)
	}

	//push
	s.stack[s.stackPointer] = val
	s.stackPointer += 1
}

func (s *DynamicStack) pop() int {

	if s.stackPointer < s.stackCapacity/3 {
		//deallocate last third
		s.stackCapacity = 2 * s.stackCapacity / 3

		var old_stack = s.stack
		s.stack = make([]int, s.stackCapacity)

		copy(s.stack, old_stack)
	}

	if s.stackPointer == 0 {
		panic("popped from empty stack")
	} else {
		s.stackPointer -= 1
		return s.stack[s.stackPointer]
	}
}

type calculator struct {
	expression         []item
	instructionPointer int
	stack              stack
}

func makeCalculator(expression []item, dynamic_stack bool) calculator {
	var stack stack
	if dynamic_stack {
		stack = newDynamicStack(4)
	} else {
		stack = newStaticStack(4)
	}
	return calculator{
		expression,
		0,
		stack,
	}
}

func (c *calculator) step() {

	var item = c.expression[c.instructionPointer]

	switch item.itemType {
	case Add:
		var second = c.stack.pop()
		var first = c.stack.pop()
		c.stack.push(first + second)
	case Sub:
		var second = c.stack.pop()
		var first = c.stack.pop()
		c.stack.push(first - second)
	case Mul:
		var second = c.stack.pop()
		var first = c.stack.pop()
		c.stack.push(first * second)
	case Div:
		var second = c.stack.pop()
		var first = c.stack.pop()
		c.stack.push(first / second)
	case Mod:
		var second = c.stack.pop()
		var first = c.stack.pop()
		c.stack.push(first % second)
	case SpecialMult:
		var second = c.stack.pop()
		var first = c.stack.pop()
		var mult = first * second

		if 0 > mult {
			panic("negative arg to special mult not defined")
		}
		if 99 < mult {
			panic("too large arg to special mult not defined")
		}

		var digit10 = mult / 10
		var digit1 = mult - digit10
		c.stack.push(digit10 + digit1)

	case Value:
		c.stack.push(item.value)
	}

	c.instructionPointer++
}

func (c *calculator) run() int {
	for c.instructionPointer < len(c.expression) {
		c.step()
	}
	return c.stack.pop()
}

func benchStacks(n int, tries int) (static_push []time.Duration, static_pop []time.Duration, dynamic_push []time.Duration, dynamic_pop []time.Duration) {
	var static = newStaticStack(n)
	var dynamic = newDynamicStack(1)

	static_push = make([]time.Duration, tries)
	static_pop = make([]time.Duration, tries)
	dynamic_push = make([]time.Duration, tries)
	dynamic_pop = make([]time.Duration, tries)

	for try := 0; try < tries; try++ {
		println("try: ", try+1)
		time_start := time.Now()
		for i := 0; i < n; i++ {
			static.push(42)
		}
		static_push[try] = time.Now().Sub((time_start))

		time_start = time.Now()
		for i := 0; i < n; i++ {
			static.pop()
		}
		static_pop[try] = time.Now().Sub((time_start))

		time_start = time.Now()
		for i := 0; i < n; i++ {
			dynamic.push(42)
		}
		dynamic_push[try] = time.Now().Sub((time_start))

		time_start = time.Now()
		for i := 0; i < n; i++ {
			dynamic.pop()
		}
		dynamic_pop[try] = time.Now().Sub((time_start))

	}

	return

}

func main() {

	var calc = makeCalculator([]item{
		item{Value, 2},
		item{Value, 9},
		item{SpecialMult, 0},

		item{Value, 1},
		item{Value, 9},
		item{SpecialMult, 0},

		item{Value, 2},
		item{Value, 0},
		item{SpecialMult, 0},

		item{Value, 1},
		item{Value, 1},
		item{SpecialMult, 0},

		item{Value, 2},
		item{Value, 0},
		item{SpecialMult, 0},

		item{Value, 1},
		item{Value, 9},
		item{SpecialMult, 0},

		item{Value, 2},
		item{Value, 4},
		item{SpecialMult, 0},

		item{Value, 1},
		item{Value, 6},
		item{SpecialMult, 0},

		item{Value, 2},
		item{Value, 7},
		item{SpecialMult, 0},

		item{Add, 0},
		item{Add, 0},
		item{Add, 0},
		item{Add, 0},
		item{Add, 0},
		item{Add, 0},
		item{Add, 0},
		item{Add, 0},

		item{Value, 10},
		item{Mod, 0},
	}, true)

	fmt.Println(calc.run())

	var a, b, c, d = benchStacks(100000000, 10)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//TODO
// error handle
//stacks
//calculator

//benchmark

//special operation
//mod
//my number
