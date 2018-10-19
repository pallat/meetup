package main

import "fmt"

func main() {
	counter := counterFactory()
	for i := 0; i < 10; i++ {
		fmt.Println(counter())
	}
}

func counterFactory() func() int {
	i := -1
	return func() int {
		i++
		return i
	}
}
