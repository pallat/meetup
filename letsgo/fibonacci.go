package main

import "fmt"

func main() {
	chNo := make(chan int)
	chQ := make(chan struct{})
	go fibonacci(chNo, chQ)
	for i := 0; i < 10; i++ {
		fmt.Println(<-chNo)
	}
	chQ <- struct{}{}
}

func fibonacci(chNo chan int, chQ chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case chNo <- a:
			a, b = b, a+b
		case <-chQ:
			return
		}
	}
}
