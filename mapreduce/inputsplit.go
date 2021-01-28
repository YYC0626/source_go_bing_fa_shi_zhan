package mapreduce

import (
	"strings"
)

func InputSplit(txt string, out [4]chan<- string) {
	txt = strings.Replace(txt, "\n", ",", -1)
	txt = strings.Replace(txt, ".", ",", -1)
	txt = strings.Replace(txt, "!", ",", -1)

	arr := strings.Split(txt, ",")
	mod := len(arr) % 4
	for i := 0; i < (4 - mod); i++ {
		arr = append(arr, "")
	}
	lenSplit := len(arr) / 4
	for i := 0; i < 4; i++ {
		go func(ch chan<- string, lines []string) {
			for _, line := range lines {
				word := strings.Split(line, " ")
				for _, w := range word {
					ch <- w
				}
			}
			close(ch)
		}(out[i], arr[i*lenSplit:(i+1)*lenSplit])
	}
}
