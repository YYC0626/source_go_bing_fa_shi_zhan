package main

import (
	"fmt"
	"time"
)

func generate1(ch chan<- int) {
	time.Sleep(3 * time.Second)
	ch <- 7
}

func generate2(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 8
}

func main() {

	ch := make(chan int)
	go generate1(ch)

	ch2 := make(chan int)
	go generate2(ch)

	for {
		select {
		case n1 := <-ch:
			fmt.Println(n1)
		case n2 := <-ch2:
			fmt.Println(n2)
		case <-time.After(time.Second):
			fmt.Println("timeout")
			goto EXIT
		}
	}

EXIT:
	var input string
	fmt.Scanln(&input)
}
