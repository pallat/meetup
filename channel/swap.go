package channel

const max = 3

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
