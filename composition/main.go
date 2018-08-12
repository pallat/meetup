package main

import (
	"fmt"
)

// ref https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=554s

type A struct {
	name string
}

func (a A) Legs() int { return 4 }
func (a A) PrintLegs() {
	fmt.Println(a.Legs())
}

func (a A) String() string {
	return "a.name"
}

func (a A) PrintName() {
	fmt.Println(a)
}

type B struct {
	A
}

func (b B) String() string {
	return b.name
}

func (b B) Legs() int {
	return 5
}

func main() {
	var a A
	var b B

	a.name = "aaaa"
	b.name = "bbbb"

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(b.Legs())
	b.PrintLegs()

	a.PrintName()
	b.PrintName()
}
