package channel_test

import (
	"fmt"
	"time"
)

func Waiter(ch chan int) {
	<-ch
	fmt.Println("I have got signal")
}

func ExampleWaiter() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go Waiter(ch)
	}

	fmt.Println("waiting for 3s")

	<-time.After(3 * time.Second)

	close(ch)

	<-time.After(time.Second)
}
