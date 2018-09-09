package main

import "fmt"

func main() {
	fn := PoolFunc(2)

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}

func PoolFunc(max int) func() int {
	ch := make(chan int, max)

	for i := 0; i < max; i++ {
		ch <- i
	}

	return func() int {
		i := <-ch
		ch <- i
		return i
	}
}
