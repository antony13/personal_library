package main

import (
	"fmt"
)

func main() {
	dict := make(map[string]string) //similar dictionary
	//dict := map[string]string
	dict["go"] = "Golang"
	dict["C#"] = "CSharp"
	dict["php"] = "PHP"

	for v := range dict {
		fmt.Printf("Key: The value is: %s\n", v)
	}
} //end of main
