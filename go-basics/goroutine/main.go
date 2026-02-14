package main

import (
	"context"
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
	// Unbuffered channel: A channel that does not have a buffer and requires both sender and receiver to be ready for communication to occur.

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

func bufferedChannel() {
	// Buffered channel: A channel that has a buffer and allows sending and receiving of values without blocking until the buffer is full or empty.

	ch := make(chan int, 3) // Create a buffered channel with a capacity of 3

	ch <- 1 // Send value to channel (non-blocking)
	ch <- 2
	ch <- 3

	close(ch)

	fmt.Println("Finished sending values to channel")

	for value := range ch { // Using when channel is closed, it will exit the loop
		fmt.Println("Received value:", value)
	}
}

func taskChannel(taskNumber int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done() // Defer ensures that Done is called when the function exits

	fmt.Println("Task", taskNumber, "is running")
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task %d is sending result to channel", taskNumber)
	ch <- fmt.Sprintf("Task %d completed", taskNumber)
}

func selectChannel() {
	// Select statement: A control structure that allows a goroutine to wait on multiple communication operations (channels) simultaneously and proceed with the one that is ready first.

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// fmt.Println(<-ch1) // This will block until a message is received from ch1
	// fmt.Println(<-ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}

func contextExample(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, exiting goroutine")
			return
		default:
			fmt.Println("Running:", ctx.Value("priority"))
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	//demoDefer()

	// fmt.Println("--------------------------------")

	// start := time.Now()

	// var wg sync.WaitGroup

	// for i := range 4 {
	// 	wg.Add(1)
	// 	go task(i+1, &wg)
	// }

	// wg.Wait()

	// fmt.Print("Total time: ", time.Since((start)))

	// fmt.Println("--------------------------------")

	// go unbufferedChannel()

	// go bufferedChannel()

	//time.Sleep(1 * time.Second) // Sleep to allow goroutines to finish before main exits

	// fmt.Println("--------------------------------")

	// start := time.Now()

	// var wg sync.WaitGroup

	// ch := make(chan string, 8) // buffered channel to hold results from goroutines
	// ch := make(chan string) // unbuffered channel to hold results from goroutines

	// for i := range 4 {
	// 	wg.Add(1)
	// 	go taskChannel(i+1, &wg, ch)
	// }

	// go func() {
	// 	// Wait for all goroutines to finish and then close the channel
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(<-ch)
	// }

	// for result := range ch {
	// 	fmt.Println(result)
	// }

	// fmt.Println("Total time:", time.Since(start))

	// fmt.Println("--------------------------------")

	// selectChannel()

	// fmt.Println("--------------------------------")

	// start := time.Now()
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	// ctx = context.WithValue(ctx, "priority", "high")

	// wg.Add(1)

	// go contextExample(ctx, &wg)

	// wg.Wait()
	// fmt.Println("Total time:", time.Since(start))

	// fmt.Println("--------------------------------")

	// burgerCh := make(chan string)
	// pizzaCh := make(chan string)

	// ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	// defer cancel()

	// go cookBurger(burgerCh, ctx)
	// go cookPizza(pizzaCh, ctx)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case burgerResult, ok := <-burgerCh:
	// 		if ok {
	// 			fmt.Println("Burger result:", burgerResult)
	// 		}
	// 	case pizzaResult, ok := <-pizzaCh:
	// 		if ok {
	// 			fmt.Println("Pizza result:", pizzaResult)
	// 		}
	// 	case <-ctx.Done():
	// 		fmt.Println("Context timeout, canceling cooking")
	// 		return
	// 	}
	// }

	fmt.Println("--------------------------------")
	raceCondition()
}

func raceCondition() {
	token := 0
	var wg sync.WaitGroup

	var mu sync.Mutex // Mutex to protect access to the shared variable

	start := time.Now()

	for range 1000 {
		wg.Add(1)
		go func() {
			mu.Lock()
			token++
			mu.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final token value:", token, "Time taken:", time.Since(start))

	// To avoid race condition, we can use sync.Mutex to lock the critical section of code that modifies the shared variable (token in this case) to ensure that only one goroutine can access it at a time.
}

func cookBurger(burgerCh chan<- string, ctx context.Context) {
	fmt.Println("Start to cook burger")
	select {
	case <-time.After(1 * time.Second):
		burgerCh <- "Burger is cooked"
	case <-ctx.Done():
		fmt.Println("Cooking burger is canceled")
		return
	}
}

func cookPizza(pizzaCh chan<- string, ctx context.Context) {
	fmt.Println("Start to cook pizza")
	select {
	case <-time.After(2 * time.Second):
		pizzaCh <- "Pizza is cooked"
	case <-ctx.Done():
		fmt.Println("Cooking pizza is canceled")
		return
	}
}
