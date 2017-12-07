package main

import (
	"fmt"
	"calc" //new lib name
)

func main() {
	var x,y int = 10,3 
	fmt.Println(calc.Add(x,y))
	fmt.Println(calc.Subtract(x,y))
}
