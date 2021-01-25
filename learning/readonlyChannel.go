package main

import (
	"fmt"
)

//write-only
func generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

//read-only
func square(ch <-chan int) {
	for {
		n := <-ch
		fmt.Printf("square(%d)=%d\n", n, n*n)
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	go square(ch)
	var input string
	fmt.Scanln(&input)
}
