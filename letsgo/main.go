package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go delay(i, time.Duration(r.Intn(5))*time.Second)
	}
	wg.Wait()

	fmt.Println(time.Now().Sub(start))
}

func delay(i int, d time.Duration) {
	time.Sleep(d)
	fmt.Println(i, d)
	wg.Done()
}
