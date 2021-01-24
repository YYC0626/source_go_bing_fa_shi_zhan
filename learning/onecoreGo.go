package main

import (
	"fmt"
	"runtime"
	"time"

	"source_go_bing_fa_shi_zhan/onecore"
)

func main() {
	//set core to 1
	runtime.GOMAXPROCS(1)

	gopher1 := onecore.Gopher{Name: "Gopher1", Id: 1}
	gopher2 := onecore.Gopher{Name: "Gopher2", Id: 2}

	go gopher1.MakeCoffee("A")
	go gopher2.MakeCoffee("B")

	time.Sleep(40 * time.Second)
	fmt.Println("==============END==============")
}
