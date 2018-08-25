package cast

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
