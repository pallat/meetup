package main

import "fmt"

func NewFibonacci() fibonacci {
	return fibonacci{a: 0, b: 1}
}

type fibonacci struct {
	a, b int
}

func (f *fibonacci) value() int {
	defer f.next()
	return f.a
}

func (f *fibonacci) next() {
	f.a, f.b = f.b, f.a+f.b
}

func main() {
	f := NewFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f.value())
	}
}

// END
