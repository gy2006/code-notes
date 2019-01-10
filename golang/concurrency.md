# Concurrency

## goroutines (lightweight thread)

```go
func startGoroutins() {
	go func() {
        for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			log.Println("i = ", i)
		}
	}()
}
```

```go 
func TestShouldStartGoroutines(t *testing.T) {
	startGoroutins()
	time.Sleep(15 * time.Second)
}
```

## wait && channel

send and receive values: `channel := make(chan int)`

```go
func startChannelDemo() {
	channel := make(chan int)

	var wt sync.WaitGroup
	wt.Add(2)

	// create producer
	producer := func() {
		defer close(channel)
		defer wt.Done()

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			log.Println("produce i = ", i)

			// wait for consumer
			channel <- i
		}
	}

	// create consumer
	consumer := func() {
		defer wt.Done()

		for {
			// wait for producer
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
```

buffered channle `channel := make(chan int, 1000)`

## select

`select` statement lets a goroutine 'wait' on multiple communication operations.

- 'blocks' until one of its cases can run
- 'random' choose one if multiple case are ready

```go
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
```

## lock & unlock