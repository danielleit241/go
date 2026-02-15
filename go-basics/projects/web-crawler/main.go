package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10000
	maxWorkers := 5
	queueCh := make(chan int, n)

	for i := range n {
		queueCh <- i
	}

	close(queueCh)

	var wg sync.WaitGroup
	count := 0
	for i := range maxWorkers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range queueCh {
				fmt.Printf("Worker %d processing item %d\n", i, j)
				count++
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("Total items processed: %d\n", count)
}
