package main

import (
	"fmt"
	"time"
)

func main() {
	// Select con timeout pattern
	c := make(chan bool)

	go func() {
		fmt.Println("Connecting to database...")
		time.Sleep(2 * time.Second)
		c <- true
	}()

	select {
	case conn := <-c:
		fmt.Println("Connected:", conn)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout...")
	}
}
