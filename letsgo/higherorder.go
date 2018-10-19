package main

import "fmt"

func myString() string {
	return "test..."
}

func main() {
	executer(myString)
}

func executer(fn func() string) {
	fmt.Println(fn())
}
