package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup // wait for the program to finish goroutines

func main() {
	wg.Add(2) // add a count of 2, one for each subroutine
	fmt.Println("Goroutine is starting...")
	go printCounts("A") //goroutine
	go printCounts("B") //goroutine
	fmt.Println("Wait to exit...")
	wg.Wait()
	fmt.Println("OK finished")

}//end of main

func printCounts(label string) {
		defer wg.Done() // wait for WaitGroup to finish and then execute it

		for count := 1; count <= 10; count++ { //randomly wait
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep)*time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}//end of printCounts func
