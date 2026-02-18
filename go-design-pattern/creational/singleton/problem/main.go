package main

import (
	"fmt"
	"sync"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool {
	return c.logAllowed
}

func NewConfig(logAllowed bool) *config {
	return &config{
		logAllowed: logAllowed,
	}
}

func main() {
	//Demo 1k request at the same time, we will create 1k config instance, which is not efficient.

	rps := 1000
	var wg sync.WaitGroup
	wg.Add(rps)

	config := NewConfig(true)

	for i := range rps {
		go func(idx int) {
			defer wg.Done()
			requestHandler(idx)

			if config.LogAllowed() {
				fmt.Printf("Request %d handled successfully.\n", idx)
			}
		}(i)
	}

	wg.Wait()
}

func requestHandler(idx int) {
	// I have some log to print here
	// I have to know the config that if it was allowed print log.
	// but I cannot modify definition of this method.
}
