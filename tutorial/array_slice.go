package main

import (
	"fmt"
)

func main() {
	x := [5]int{0, 1, 2, 2, 3} //array
	fmt.Println(x)
	//y:= []int{} //slice: something like array but without predefined length
	y := []int{} //better to declare it like this
	y1 := append(y, 13)
	fmt.Println(y1)
	fmt.Printf("Length of array %d and its capacity %d\n", cap(x), len(x))
	fmt.Printf("Lenght of slice %d and its capacity %d\n", cap(y1), len(y1))

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}          //array
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 9} //slice
	fmt.Println("\nArray:")
	for i, v := range a {
		fmt.Printf("Index %d and Value: %d\n", i, v)
	} //end for

	fmt.Println("\nSlice:")
	for i, v := range b {
		fmt.Printf("Index %d and Value: %d\n", i, v)
	} //end for

} //end of main function
