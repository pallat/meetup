package cast_test

import (
	"fmt"
)

type (
	Int    int
	String string
)

func ExampleSimple() {
	var n Int = 24

	var i int = int(n)
	fmt.Printf("%T %v", n, n)
	fmt.Printf("%T %v", i, i)
}
