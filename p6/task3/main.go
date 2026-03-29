package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func startServer(ctx context.Context, name string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
				out <- fmt.Sprintf("[%s] metric: %d", name, rand.Intn(100))
			}
		}
	}()

	return out
}

func FanIn(ctx context.Context, ctotal ...<-chan string) <-chan string {
	result := make(chan string)
	var wg sync.WaitGroup

	for _, c := range ctotal {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				result <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(result)
		fmt.Println("closed channel")
	}()

	go func() {
		<-ctx.Done()
		_, ok := <-result
		if ok {
			close(result)
			fmt.Println("close fanin")
		}
		fmt.Println(ctx.Err(), "fanin canceled")
	}()

	return result
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch1 := startServer(ctx, "Alpha")
	ch2 := startServer(ctx, "Beta")
	ch3 := startServer(ctx, "Gamma")
	ch4 := FanIn(ctx, ch1, ch2, ch3)

	for val := range ch4 {
		fmt.Println(val)
	}
}
