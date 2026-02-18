package main

import (
	"log"
	"sync"
	"time"
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

var singletonConfig = &application{once: sync.Once{}}

func GetApplication() *application {
	return singletonConfig
}

type application struct {
	once sync.Once
	cfg  *config
}

func (app *application) GetConfig() *config {
	// This is not thread-safe, and it will create multiple config instances if called concurrently.
	// if app.cfg == nil {
	// 	log.Println("It should be run only once. But you'll it many times!!")
	// 	app.loadConfig()
	// }

	// return app.cfg

	if app.cfg == nil {
		app.once.Do(func() {
			log.Println("Loading config once and forever.")
			app.loadConfig()
		})
	}

	return app.cfg
}

func (app *application) loadConfig() {
	time.Sleep(100 * time.Millisecond) // Simulate time-consuming config loading
	app.cfg = &config{logAllowed: true}
}

func main() {
	//Demo 1k request at the same time, we will create 1k config instance, which is not efficient.

	rps := 1000
	var wg sync.WaitGroup
	wg.Add(rps)

	for i := range rps {
		go func(idx int) {
			defer wg.Done()
			requestHandler(idx)

			if GetApplication().GetConfig().LogAllowed() {
				log.Printf("Request %d handled successfully.\n", idx)
			}
		}(i)
	}

	wg.Wait()
}

func requestHandler(idx int) {
	if GetApplication().GetConfig().LogAllowed() {
		log.Printf("Handling request %d... please wait.\n", idx)
	}
}
