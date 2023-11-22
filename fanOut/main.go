package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case m, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Worker %d received %d\n", id, m)
		}
	}
}

func fanOut(in <-chan int, numWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, in, &wg)
	}
	wg.Wait()
}

func main() {
	ch := make(chan int)

	go fanOut(ch, 3)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)

	time.Sleep(time.Second * 5)
}
