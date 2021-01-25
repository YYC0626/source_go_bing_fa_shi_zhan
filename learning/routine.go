package main

import (
	"fmt"
	"time"
)

// 1
//
// func main() {
//     go printHello("Hello Goroutine")
//     time.Sleep(1 * time.Second)
// }
//
// func printHello(s string) {
//     fmt.Println(s)
// }

//2
//
// func main() {
//     go func(s string) {
//         fmt.Println(s)
//     }("Hello Goroutine")
//     time.Sleep(1 * time.Second)
// }

func main() {

	printHello := func(s string) {
		fmt.Println("Hello Goroutine")
	}
	go printHello("hello Goroutine")

	time.Sleep(time.Second * 1)
}
