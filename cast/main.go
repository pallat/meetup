package main

import (
	"fmt"
)

type A struct {
	Foo string
	Bar string
	Baz string
}

func (a A) String() string {
	return a.Foo + a.Bar + a.Baz
}

type B struct {
	Foo string
	Bar string
	Baz string
}

func (b B) String() string {
	return b.Baz + b.Bar + b.Foo
}

func main() {
	a := A{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}

	fmt.Println(a)

	b := B{
		Foo: "fooooo",
		Bar: "barrrr",
		Baz: "bazzzz",
	}

	fmt.Println(b)

	b = B(a)

	fmt.Println(b)
}
