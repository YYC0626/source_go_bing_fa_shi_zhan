package main

import (
	"fmt"

	"source_go_bing_fa_shi_zhan/calc"
)

func main() {
	p, n, e := 0, 0, 0

	fmt.Scanln(&p, &n, &e)

	Z, A, c := calc.AtomCalc(p, n, e)

	fmt.Println("while p=%d, n=%d, e=%d\n", p, n, e)
	fmt.Println("Z = %d, A = %d, c = %d\n", Z, A, c)
}
