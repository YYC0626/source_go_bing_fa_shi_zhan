package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter = 0
	func() {
		for i := 0; i < 1000; i++ {
			go func() {
				counter++
			}()
		}
	}()

	fmt.Println("over!press any key to exit.")

	var input string
	fmt.Scanln(&input)
	fmt.Println(counter)

	//1-2个cpu时cnt为1000
}
