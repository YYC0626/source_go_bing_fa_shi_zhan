package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter = 0
	var mlock sync.Mutex

	func() {
		for i := 0; i < 1000; i++ {
			go func() {
				mlock.Lock()
				counter++
				mlock.Unlock()
			}()
		}
	}()

	fmt.Print("\nover.press any key to exit...")
	var input string
	fmt.Scanln(&input)
	fmt.Println(counter)
}
