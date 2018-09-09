package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan time.Time)
	go sender(ch)
	receiver(ch)
}

func sender(ch chan<- time.Time) {
	fmt.Println("sender: waiting for receiver")
	ch <- time.Now()
	fmt.Println("sender: already sent")
}

func receiver(ch <-chan time.Time) {
	fmt.Println("receiver: waiting for 5 seconds")
	time.Sleep(time.Second * 5)
	fmt.Println("receiver: ready")
	fmt.Println(time.Since(<-ch))
}
