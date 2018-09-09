package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan time.Time, 1)
	go sender(ch)
	receiver(ch)
}

func sender(ch chan<- time.Time) {
	fmt.Println("sender: reserved for data")
	ch <- time.Now()
	fmt.Println("sender: already sent to buffer")
}

func receiver(ch <-chan time.Time) {
	fmt.Println("receiver: waiting for 5 seconds")
	time.Sleep(time.Second * 5)
	fmt.Println("receiver: ready")
	fmt.Println(time.Since(<-ch))
}
