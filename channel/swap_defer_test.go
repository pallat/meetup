package channel_test

import "fmt"

const max = 3

func SwapFuncDefer() func() int {
	ch := make(chan int, max)

	for i := 0; i < max; i++ {
		ch <- i
	}

	return func() int {
		i := <-ch
		defer func() { ch <- i }()
		return i
	}
}

func ExampleSwapFunc_defer() {
	fn := SwapFuncDefer()

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
