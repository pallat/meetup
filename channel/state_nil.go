package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var ch chan string

	go func() {
		println(<-ch)
		wg.Done()
	}()

	go func() {
		ch <- "go"
	}()

	ch = make(chan string)
	wg.Wait()
}
