package main

import (
	"fmt"
	"sync"
)

func withSyncMap() {
	var safeMap sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			safeMap.Store("key", key)
		}(i)
	}

	wg.Wait()
	value, _ := safeMap.Load("key")
	fmt.Println("Using sync.Map:", value)
}

func withRWMutex() {
	var rwm sync.RWMutex
	var wg sync.WaitGroup
	unsafeMap := make(map[string]int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			rwm.Lock()
			unsafeMap["key"] = key
			rwm.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Using sync.RWMutex:", unsafeMap["key"])
}

func main() {
	withSyncMap()
	withRWMutex()
}
