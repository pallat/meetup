package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 8
	}()
	fmt.Println(<-ch)
}
