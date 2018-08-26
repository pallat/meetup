package cast_test

import "fmt"

type Swapper interface {
	Do(int, int) (int, int)
}

type Swap func(int, int) (int, int)

func (s Swap) Do(a, b int) (int, int) {
	return s(a, b)
}

func Print(s Swapper, a, b int) {
	fmt.Println(s.Do(a, b))
}

func ExampleFunc() {
	var swap = func(a, b int) (int, int) {
		return b, a
	}

	Print(Swap(swap), 0, 1)
}
