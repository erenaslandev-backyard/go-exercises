package main

import (
	"fmt"
	"sync"
	"time"
)

func runChannel() {
	fmt.Println("Begin.")
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Schedule %d\n", i)
			channel <- i
		}
		close(channel)
	}()

	for i := range channel {
		fmt.Println(i)
	}

	fmt.Println("Done!")
}

func runChannelWithDone() {
	fmt.Println("Begin.")
	done := make(chan interface{})
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Schedule %d\n", i)
			channel <- i
		}
		close(channel)
		close(done)
	}()

	for i := range channel {
		fmt.Println(i)
	}

	<-done
	fmt.Println("Done!")
}

func runReadOnlyChannel() {
	fmt.Println("Begin.")
	// channel := make(chan int)
	channel := make(chan int, 10)

	numberOfTasks := 10
	var wg sync.WaitGroup
	wg.Add(1)

	go func(wChannel chan<- int, n int) {
		for i := 0; i < n; i++ {
			wg.Add(1)
			fmt.Printf("Schedule %d\n", i)
			wChannel <- i
		}
		close(wChannel)
		wg.Done()
	}(channel, numberOfTasks)

	go func(rChannel <-chan int) {
		for i := range rChannel {
			fmt.Printf("Start %d\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Finish %d\n", i)
			wg.Done()
		}
		// close(rChannel) Cannot close receive only channel
	}(channel)

	// <-channel
	wg.Wait()
	fmt.Println("Done!")
}

func sendToNilChannel() {
	var c chan int
	// c := make(chan int)
	c <- 1
}

func closeClosedChannel() {
	c := make(chan struct{})
	close(c)
	close(c)
}

func sendClosedChannel() {
	c := make(chan struct{})
	close(c)
	c <- struct{}{}
}

func receiveFromClosedChannel() {
	// c := make(chan struct{})
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	_, open := <-c
	fmt.Println("open?:", open)
}

func selectFromClosedChannel() {
	// c := make(chan struct{})
	c := make(chan int)
	close(c)

	for {
		select {
		case val := <-c:
			fmt.Println(val)
		case <-time.After(10 * time.Millisecond):
			fmt.Println("timeout")
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	// closeClosedChannel()
	// sendToNilChannel()
	// sendClosedChannel()
	// receiveFromClosedChannel()
	// selectFromClosedChannel()

	// runChannelWithDone()
	// runChannel()
	// runReadOnlyChannel()

	go func() { select {} }()
	select {}
}
