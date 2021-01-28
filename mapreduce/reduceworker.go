package mapreduce

func ReduceCountWorker(in <-chan map[string]int, out chan<- map[string]int) {
	count := map[string]int{}
	for n := range in {
		for word := range n {
			count[word] = count[word] + n[word]
		}
	}
	out <- count
	close(out)
}
