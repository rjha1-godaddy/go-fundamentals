package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printMessage("Hello from goroutine", &wg)

	wg.Wait()
	fmt.Println("All goroutines completed!")
}
