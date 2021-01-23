package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("auto EXIT\n")
	f := bufio.NewReader(os.Stdin)
	input := ""
	str := ""
	for {
		fmt.Print(">")
		input, _ = f.ReadString('\n')
		if len(input) == 1 {
			continue
		}
		fmt.Sscan(input, &str) //移除换行符
		if str == "exit" {
			break
		} else {
			fmt.Printf("input:%s\n", str)
		}
	}

}
