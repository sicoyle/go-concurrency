package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker is our goroutine that prints what worker has started/stopped
// Note: WaitGroups must be passed to functions by pointer
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0
	wg.Wait()
}
