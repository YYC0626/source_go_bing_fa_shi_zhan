package calc

import (
	"fmt"
)

const title = "原子计算器v1.0，输入质子、中子、电子的值，可以得到原子序数、质量和电子"
const usage = "16 16 18"

func init() {
	fmt.Println(title)
	fmt.Println(usage)
	fmt.Println("=======================")
}

func AtomCalc(p, n, e int) (int, int, int) {

	Z := p

	A := p + n

	c := p - e

	return Z, A, c
}
