package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	go busy(quit)
	go tick()

	fmt.Println("Waiting for signal")
	<-quit
	fmt.Println("Finish")
}

func busy(quit chan struct{}) {
	time.Sleep(5 * time.Second)
	quit <- struct{}{}
}

func tick() {
	for range time.Tick(time.Second) {
		fmt.Println(".")
	}
}
