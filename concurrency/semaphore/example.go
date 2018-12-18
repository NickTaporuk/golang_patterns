package semaphore

import (
	"log"
	"time"
)

// main is the entry point for the application
func Example() {
	log.Println("Starting Process")

	// Create a new readerWriter with a max of 3 reads at a time
	// and a total of 6 reader goroutines.
	first := start("First", 3, 6)

	// Create a new readerWriter with a max of 1 reads at a time
	// and a total of 1 reader goroutines.
	second := start("Second", 2, 2)

	// Let the program run for 2 seconds.
	time.Sleep(2 * time.Second)

	// Shutdown both of the readerWriter processes.
	shutdown(first, second)

	log.Println("Process Ended")
	return
}
