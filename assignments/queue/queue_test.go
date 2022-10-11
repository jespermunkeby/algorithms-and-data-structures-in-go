package main

import (
	"math/rand"
	"testing"
)

func TestStaticQueue(t *testing.T) {
	q := makeStaticQueue(500)

	//some statistical tests that as a sort of catch all

	for try := 0; try < 1000; try++ {
		maxN := rand.Intn(400)
		for n := 0; n < maxN; n++ {
			q.add(n)
		}

		for n := 0; n < maxN; n++ {
			if v, ok := q.remove(); ok {
				if v != n {
					panic("static stack doesent work!")
				}
			}
		}
	}

}

func TestDynamicQueue(t *testing.T) {
	q := makeDynamicQueue()

	//some statistical tests that as a sort of catch all

	for try := 0; try < 1000; try++ {
		maxN := rand.Intn(400)
		for n := 0; n < maxN; n++ {
			q.add(n)
		}

		for n := 0; n < maxN; n++ {
			if v, ok := q.remove(); ok {
				if v != n {
					panic("static stack doesent work!")
				}
			}
		}
	}

}
