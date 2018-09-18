package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}

func sender(ch chan<- int) {
	fmt.Println("sender: waiting for receiver")
	ch <- 8
	fmt.Println("sender: already sent")
}

func receiver(ch <-chan int) {
	fmt.Println("receiver: waiting for 5 seconds")
	time.Sleep(time.Second * 5)
	fmt.Println("receiver: ready")
	fmt.Println("get", <-ch)
}
