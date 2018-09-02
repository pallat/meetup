package channel_test

import (
	"fmt"
)

const max = 3

func SwapFunc() func() int {
	ch := make(chan int, max)

	for i := 0; i < max; i++ {
		ch <- i
	}

	turn := make(chan int)

	go func() {
		for c := range turn {
			ch <- c
		}
	}()

	return func() int {
		i := <-ch
		turn <- i
		return i
	}
}

func ExampleSwapFunc_chan() {
	fn := SwapFunc()

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
