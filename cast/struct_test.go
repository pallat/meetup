package cast_test

import (
	"fmt"
)

type Struct struct {
	Foo string
	Bar string
	Baz string
}

func (s Struct) String() string {
	return s.Foo + s.Bar + s.Baz
}

type StructLike struct {
	Foo string
	Bar string
	Baz string
}

func (s StructLike) String() string {
	return s.Baz + s.Bar + s.Foo
}

func ExampleStruct() {
	a := Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}

	fmt.Println(a)

	b := StructLike{
		Foo: "fooooo",
		Bar: "barrrr",
		Baz: "bazzzz",
	}

	fmt.Println(b)

	b = StructLike(a)

	fmt.Println(b)
}
