package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func printCounts(label string) {
	// Schedule the call to WaitGroup's Done tell are done.
	defer wg.Done()

	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Count: %d from %s\n", count, label)
	}
}

func main() {
	// Add a count of two, one for each goroutines
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// launch a goroutine with label "A"
	go printCounts("A")
	// launch a goroutine with label "B"
	go printCounts("B")
	// Wait for the goroutines to finish.
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
