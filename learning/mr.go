package main

import (
	"fmt"
	"time"

	"source_go_bing_fa_shi_zhan/mapreduce"
)

func main() {
	t := time.Now()
	input := "Go is awsome, Go is Best, Best are good!\nYou are Hero, Thank You."
	fmt.Println(mapreduce.Run(input))
	elapsed := time.Since(t)
	fmt.Println("time elasped:", elapsed)

	// $ go run mr.go -race
}
