package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // wait for the program to finish with goroutines before exit

func main() {
	count := make(chan int)
	wg.Add(2)

	fmt.Println("Starting go routines!")
	go printCounts("A", count)
	go printCounts("B", count)
	fmt.Println("Channel is starting...")
	count <- 1 //we send 1 to the channel
	fmt.Println("Waiting to finish!")
	wg.Wait()
	fmt.Println("Exiting application")
} //end of main

//Takes as an input the label ('A' or 'B') and the channel integer
func printCounts(label string, count chan int) {
	defer wg.Done() //Let the execution know that we finished.
	for {
		val, msg := <-count
		if !msg {
			fmt.Println("Channel was closed, no more inputs found\n")
			return
		} //end if

		fmt.Printf("Count %d received from %s\n", val, label)
		if val == 10 {
			fmt.Printf("Channel is closed from %s\n", label)
			close(count) //close the channel
			return
		} //end if

		val++        //increment the val
		count <- val //Send count to the other routine
	} // end for
} //end of printCounts
