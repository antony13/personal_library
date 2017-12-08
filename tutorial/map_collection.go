package main

import (
	"fmt"
)

type CityLoc struct {
	lat, long float64
}

func main() {
	dict := make(map[string]string) //similar to python dictionary
	dict["go"] = "Golang"
	dict["C#"] = "CSharp"
	dict["php"] = "PHP"

	for v := range dict {
		fmt.Printf("Key: The value is: %s\n", v)
	}
	
	//create a map data structure with a string index and a CityLoc value
	real_map := make(map[string]CityLoc)
	real_map["athens"] = CityLoc {
		37.9838, 23.7275,
	}

} //end of main
