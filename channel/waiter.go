package channel

import (
	"fmt"
)

func Waiter(ch chan int) {
	<-ch
	fmt.Println("I have got signal")
}
