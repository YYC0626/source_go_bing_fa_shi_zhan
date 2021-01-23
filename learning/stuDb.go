package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"source_go_bing_fa_shi_zhan/ydbase"
)

func main() {
	fmt.Print("#>")

	uname := ""
	upwd := ""

	for {
		n, err := fmt.Scanln(&uname, &upwd)

		if err == io.EOF {
			break
		}

		if n > 0 {
			if !ydbase.Login(uname, upwd) {
				fmt.Println("failed to login!")
				fmt.Print("#>")
			} else {
				fmt.Println("Welcome !")
				fmt.Print("$>")
				break
			}
		}
	}

	cmd := ""
	exp := ""
	for {
		n, err := fmt.Scanln(&cmd, &exp)

		if err == io.EOF {
			break
		}

		if n > 0 {
			if cmd == "init" {

				len1, _ := strconv.Atoi(exp)
				for i := 1; i <= len1; i++ {
					stu := ydbase.Student{
						XH:     "09" + strconv.Itoa(i),
						Name:   "Name" + strconv.Itoa(i),
						Age:    22 + i,
						Height: 1.7,
					}

					ydbase.Insert(stu)
				}

				fmt.Printf("succeed.    %d lines\n", ydbase.Count())
				fmt.Print("$>")

			} else if cmd == "insert" {

				fileds := strings.Split(exp, ",")
				leng := len(fileds)

				if leng != 4 {
					fmt.Println("insert needs 4 args...")
					fmt.Print("$>")
				}
				age, _ := strconv.Atoi(fileds[2])

				height64, _ := strconv.ParseFloat(fileds[3], 32)

				height := float32(height64)

				stu := ydbase.Student{
					XH:     fileds[0],
					Name:   fileds[1],
					Age:    age,
					Height: height,
				}

				ydbase.Insert(stu)
				fmt.Printf("succeed.    %d lines.\n", ydbase.Count())
				fmt.Print("$>")

			} else if cmd == "delete" {

				ydbase.Delete(exp)
				fmt.Printf("succeed.    %d lines.\n", ydbase.Count())
				fmt.Print("$>")

			} else if cmd == "select" {

				exp := strings.Replace(exp, "%", " like ", 1)
				stu := ydbase.Query(exp)
				fmt.Printf("succeed.    %d lines.\n", len(stu))

				if len(stu) > 0 {
					ydbase.Print(stu)
				}
				fmt.Print("$>")

			} else if cmd == "update" {

				fileds := strings.Split(exp, ",")
				leng := len(fileds)

				if leng != 4 {
					fmt.Println("insert needs 4 args...")
					fmt.Print("$>")
				}

				stu := ydbase.Query("xh==" + fileds[0])
				if len(stu) == 1 {
					stu[0].Name = fileds[1]
					age, _ := strconv.Atoi(fileds[2])

					height64, _ := strconv.ParseFloat(fileds[3], 32)

					height := float32(height64)
					stu[0].Age = age
					stu[0].Height = height

					ydbase.Modify(stu[0])
					fmt.Printf("succeed. 1 line.\n")
				}
				fmt.Print("$>")

			} else if cmd == "count" {

				fmt.Printf("succeed.    %d lines.\n", ydbase.Count())
				fmt.Print("$>")

			} else if cmd == "exit" {

				break

			} else {

				fmt.Println("just support 'init', 'delete', 'update', 'select' .")
				fmt.Print("$>")

			}

		}
	}

	fmt.Print("#>")

	// root 123
	// init 10
	// select age<=25
	// count
	// delete 091
	// insert 001,jack,33,1.81
	// select name==jack
	// update 001,jack2,32,1.79
	// select name%Nam -> name like '%Nam%'

}
