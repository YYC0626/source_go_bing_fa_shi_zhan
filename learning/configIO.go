package main

import (
	"flag"
	"fmt"

	"source_go_bing_fa_shi_zhan/property"
)

func main() {
	//go run configIO.go -fkey db.uname
	content := flag.String("fkey", "db.name", "-fkey指定配置文件的key")
	flag.Parse()
	fmt.Printf("key[%s]->%s", *content, property.Read(*content))
}
