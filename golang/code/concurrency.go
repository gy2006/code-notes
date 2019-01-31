package main

import (
	"log"
	"sync"
	"time"
)

func start() {
	var wt sync.WaitGroup
	wt.Add(1)

	go func() {
		defer wt.Done()

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			log.Println("i = ", i)
		}
	}()

	wt.Wait()
}

func startChannelDemo() {
	channel := make(chan int, 10000)

	var wt sync.WaitGroup
	wt.Add(2)

	// create producer
	producer := func() {
		defer close(channel)
		defer wt.Done()

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Millisecond)
			log.Println("produce i = ", i)
			channel <- i
		}
	}

	// create consumer
	consumer := func() {
		defer wt.Done()

		for {
			i, ok := <-channel

			if !ok {
				break
			}

			log.Println("consumer i = ", i)
			time.Sleep(1 * time.Second)
		}
	}

	// start
	go producer()
	go consumer()

	// wait for down
	wt.Wait()
}

func selectDemo() {
	channel := make(chan int)
	defer close(channel)

	var wt sync.WaitGroup
	wt.Add(1)

	// create consumer only
	consumer := func() {
		defer wt.Done()

		for {
			select {
			case i, ok := <-channel:
				if !ok {
					break
				}

				log.Println("consumer i = ", i)
				time.Sleep(1 * time.Second)

			// exit after 10 seconds
			case <-time.After(10 * time.Second):
				log.Println("timeout and exit")
				return
			}
		}
	}

	go consumer()

	wt.Wait()
}
