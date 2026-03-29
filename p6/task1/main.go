package main

import (
	"fmt"
	"sync"
)

func withSyncMap() {
	var safeMap sync.Map

	for i := 0; i < 100; i++ {
		go func(key int) {
			safeMap.Store("key", key)
		}(i)
	}

	value, _ := safeMap.Load("key")
	fmt.Println("Using sync.Map:", value)
}

func main() {
	withSyncMap()
}
