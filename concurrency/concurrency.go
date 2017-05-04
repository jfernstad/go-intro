package main

import (
	"fmt"
	"math"
	"time"
)

// This should take some time to execute
func bigTask(taskID int, done chan int) {

	var res float64
	for idx := 0; idx < 800000000; idx++ {
		res += math.Cos(float64(idx) / math.Pi)
	}
	done <- taskID
}

func main() {

	// How many tasks should we run in paralell?
	concurrency := 2

	// `channels` can be used to signal between tasks/threads
	done := make(chan int, concurrency)

	// Start tasks and timer
	start := time.Now()

	for idx := 0; idx < concurrency; idx++ {
		fmt.Printf("Starting task %d\n", idx+1)

		// Magic happens here
		go bigTask(idx+1, done)
	}

	// Don't wait forever
	// This generates a timeout channel
	// that will receive something after the duration
	timeout := time.Tick(60 * time.Second)

	// Wait until all tasks have signalled that they are done, or timeout
	for c := 0; c < concurrency; c++ {

		// Select blocks until one of the `cases` return (if you have mul)
		select {
		case taskID, ok := <-done: // Read both value and result, `ok` indicated if we were closed or not
			if ok {
				fmt.Printf("%d done, %d to go\n", taskID, concurrency-c-1)
			} else {
				fmt.Printf(" Aborted \n")
			}
		case <-timeout:
			fmt.Print("Timeout !! \n")
			close(done)
		}
	}

	// Ok, they are all done now
	end := time.Now()
	fmt.Printf("Ending. Took %v\n", end.Sub(start))
}
