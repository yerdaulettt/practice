package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func withMutex() {
	var counter int
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			counter++
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Counter using mutex:", counter)
}

func withAtomic() {
	var counter atomic.Int32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println("Counter using atomic:", counter.Load())
}

func main() {
	// Teamlead, previous code result can't be always 1000, because
	// data race happens and leads to incorrect result
	withMutex()
	withAtomic()
}
