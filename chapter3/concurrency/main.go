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
