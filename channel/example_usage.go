package main

import "fmt"

func main() {
	chmsg := make(chan string)
	go ping(chmsg)
	fmt.Println(<-chmsg)
}

func ping(ch chan string) {
	ch <- "pong"
}
