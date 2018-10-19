package main

import (
	"time"
)

func main() {
	go routine("Hi")
	go routine("สวัสดี")
	time.Sleep(time.Millisecond)
	println("Go!!")
}

func routine(s string) {
	println(s)
}
