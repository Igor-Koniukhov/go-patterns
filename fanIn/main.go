package main

import (
	"fmt"
	"sync"
	"time"
)

func fanIn(ch ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	send := func(n <-chan int) {
		defer wg.Done()
		for {
			select {
			case m, ok := <-n:
				if !ok {
					return
				}
				out <- m
			}
		}
	}

	wg.Add(len(ch))
	for _, c := range ch {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generate(arr []int, dur time.Duration) <-chan int {
	ch := make(chan int)

	go func() {
		for _, n := range arr {
			ch <- n
			time.Sleep(dur)
		}
		close(ch)
	}()

	return ch
}

func main() {
	ch1 := generate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 500*time.Millisecond)
	ch2 := generate([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 300*time.Millisecond)

	m := fanIn(ch1, ch2)

	for ms := range m {
		fmt.Println(ms)
	}
}
