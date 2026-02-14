package main

import (
	"fmt"
	"sync"
	"time"
)

func task(taskNumber int, wg *sync.WaitGroup) {
	defer wg.Done() // Defer ensures that Done is called when the function exits

	println("Task", taskNumber, "is running")
	time.Sleep(1 * time.Second)
	println("Task", taskNumber, "is completed")
}

func demoDefer() {
	fmt.Println("Running 1")
	defer fmt.Println("Stop")
	fmt.Println("Running 2")
}

func unbufferedChannel() {
	ch := make(chan int)

	go func() {
		defer close(ch)                                   // Close the channel when the goroutine finishes
		ch <- 1                                           // Block1
		ch <- 2                                           // Block2
		ch <- 3                                           // Block3 -> If don't have receiver, it will block here
		fmt.Println("Finished sending values to channel") // If the channel is blocked, this line will not be executed until the channel is unblocked
	}()

	// Anonymous goroutine to receive value from channel
	// go func() {
	// 	value := <-ch
	// 	fmt.Println("Received value:", value)
	// }()

	// for i := 0; i < 3; i++ {
	// 	value := <-ch
	// 	fmt.Println("Received value:", value)
	// }

	for value := range ch { // Using when channel is closed, it will exit the loop
		fmt.Println("Received value:", value)
	}
}

func main() {
	//demoDefer()

	// start := time.Now()

	// var wg sync.WaitGroup

	// for i := range 4 {
	// 	wg.Add(1)
	// 	go task(i+1, &wg)
	// }

	// wg.Wait()

	// fmt.Print("Total time: ", time.Since((start)))

	go unbufferedChannel()

	time.Sleep(1 * time.Second) // Sleep to allow goroutines to finish before main exits
}
