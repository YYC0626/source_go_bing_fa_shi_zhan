package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func inputStream(ch chan string) {
	var i = 0
	for {
		ch <- strconv.Itoa(i) + "\n"
		time.Sleep(1 * time.Second)
		i++
	}
}

func outStream(ch chan string) {
	for {
		data := <-ch
		buf := bytes.NewBufferString(data)

		IOCopy(buf, os.Stdout)
	}
}

func IOCopy(src io.Reader, dst io.Writer) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
	}
}

func main() {
	ch := make(chan string, 2)
	go inputStream(ch)
	go outStream(ch)
	var input string
	fmt.Scanln(&input)
}
