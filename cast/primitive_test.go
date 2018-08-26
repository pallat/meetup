package cast_test

import (
	"fmt"

	"github.com/pallat/meetup/cast"
)

func ExampleInt() {
	var n cast.Int = 24

	var i int = int(n)
	fmt.Printf("%T %v", n, n)
	fmt.Printf("%T %v", i, i)
}
