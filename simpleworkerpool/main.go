package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	id int
	wg *sync.WaitGroup
}

func work(t task) {
	time.Sleep(30 * time.Microsecond)
	fmt.Printf("%+v\n", t)
	t.wg.Done()
}

func schedule(tasks <-chan task) {
	for t := range tasks {
		go work(t)
	}
}

func main() {
	tasks := make(chan task)
	go schedule(tasks)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		tasks <- task{i, &wg}
	}

	close(tasks)
	wg.Wait()
}
