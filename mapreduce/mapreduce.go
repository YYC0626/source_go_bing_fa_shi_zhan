package mapreduce

func Run(txt string) map[string]int {
	go InputSplit(txt, [4]chan<- string{chtxt1, chtxt2, chtxt3, chtxt4})

	go MapWorker(chtxt1, chmap1)
	go MapWorker(chtxt2, chmap2)
	go MapWorker(chtxt3, chmap3)
	go MapWorker(chtxt4, chmap4)

	go Master([]<-chan map[string]int{chmap1, chmap2, chmap3, chmap4}, [2]chan<- map[string]int{chreduce1, chreduce2})

	go ReduceCountWorker(chreduce1, chtemp1)
	go ReduceCountWorker(chreduce2, chtemp2)

	return SumReduce([]<-chan map[string]int{chtemp1, chtemp2})
}
