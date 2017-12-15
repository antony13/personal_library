package main

import "fmt"

func main() {
	messages := make(chan string, 2) //create a channel that accepts only 2 strings
	messages <- "Hello"
	messages <- "World"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
} //enf of main
