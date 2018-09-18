package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go sender(ch)
	go sender(ch)

	receiver(ch)
	fmt.Println("I'm going to throw away the rest")
}

func sender(ch chan<- int) {
	fmt.Println("sender: reserved for data")
	ch <- 8
	fmt.Println("sender: already sent 8 to buffer")
}

func receiver(ch <-chan int) {
	fmt.Println("receiver: waiting for 5 seconds")
	time.Sleep(time.Second * 5)
	fmt.Println("receiver: ready")
	fmt.Println("get", <-ch)
}
