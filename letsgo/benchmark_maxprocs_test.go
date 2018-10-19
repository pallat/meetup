package main

import (
	"runtime"
	"testing"
)

func BenchmarkParallel2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < b.N; i++ {
		go func() { _ = i * i }()
	}
}

func BenchmarkConcurrency(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		go func() { _ = i * i }()
	}
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
