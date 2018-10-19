package main_test

import (
	"runtime"
	"testing"
)

func BenchmarkNoRoutine(b *testing.B) {
	runtime.GOMAXPROCS(4)
	for i := 0; i < b.N; i++ {
		walkThrough()
	}
}

func walkThrough() int {
	ch := make(chan int)
	go walk(tree, ch)
	sum := 0
	for i := 0; i < 16; i++ {
		sum += <-ch
	}

	return sum
}

func walk(t *node, ch chan int) {
	ch <- t.value

	if t.left != nil {
		walk(t.left, ch)
	}
	if t.right != nil {
		walk(t.right, ch)
	}
}

type node struct {
	value int
	left  *node
	right *node
}

var tree = &node{
	value: 1,
	left: &node{
		value: 2,
		left: &node{
			value: 4,
			left: &node{
				value: 8,
			},
			right: &node{
				value: 9,
				right: &node{
					value: 10,
					left: &node{
						value: 11,
					},
				},
			},
		},
		right: &node{
			value: 5,
			right: &node{
				value: 12,
				right: &node{
					value: 13,
					left: &node{
						value: 14,
						left: &node{
							value: 15,
						},
						right: &node{
							value: 16,
						},
					},
				},
			},
		},
	},
	right: &node{
		value: 3,
		left: &node{
			value: 6,
		},
		right: &node{
			value: 7,
		},
	},
}
