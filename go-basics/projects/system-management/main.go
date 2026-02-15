package main

import (
	"context"
	"sync"

	"example.com/go/processer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go processer.RunMonitoring(ctx, &wg)

	wg.Wait()

	cancel()
}
