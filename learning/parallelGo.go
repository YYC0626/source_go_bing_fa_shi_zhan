package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"source_go_bing_fa_shi_zhan/parallel"
)

func main() {
	//set core to 2
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	gopher1 := parallel.Gopher{Name: "Gopher1", Id: 1}
	gopher2 := parallel.Gopher{Name: "Gopher2", Id: 2}

	t := time.Now()

	go gopher1.MakeCoffee("A", &wg)
	go gopher2.MakeCoffee("B", &wg)
	wg.Wait()

	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)

	fmt.Println("==============END==============")
}
