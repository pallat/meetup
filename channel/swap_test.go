package channel_test

import (
	"fmt"
	"testing"
	"time"
)

// const max = 3

func SwapFunc() func() int {
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

func ExampleSwapFunc_chan() {
	fn := SwapFunc()

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}

func BenchmarkSwap(b *testing.B) {
	fn := SwapFunc()

	for i := 0; i < b.N; i++ {
		fn()
	}
}

func TestSwapFunc_chan(t *testing.T) {
	fn := SwapFunc()

	for i := 0; i < 100; i++ {
		go fmt.Println(fn())
	}
	time.Sleep(time.Second)
}
