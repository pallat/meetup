package cast_test

import (
	"fmt"

	"github.com/pallat/meetup/cast"
)

// Output: xxx
func ExampleStruct() {
	a := cast.Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}

	fmt.Println(a)

	b := cast.StructLike{
		Foo: "fooooo",
		Bar: "barrrr",
		Baz: "bazzzz",
	}

	fmt.Println(b)

	b = cast.StructLike(a)

	fmt.Println(b)
}

func ExampleInt() {
	var n cast.Int = 24

	var i int = int(n)
	fmt.Printf("%T %v", n, n)
	fmt.Printf("%T %v", i, i)
}

func ExampleSwapper() {
	var swap = func(a, b int) (int, int) {
		return b, a
	}

	cast.Print(cast.Swap(swap), 0, 1)
}
