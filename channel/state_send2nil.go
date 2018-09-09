package main

import "fmt"

func main() {
	ch := make(chan int)

	fmt.Println(<-ch)
}
