// Run the test with go test <filename>_test.go
package calc

import (
	"testing"
	"calc"
)

func TestAdd(t *testing.T) {
	var result int
	result = calc.Add(15,10)
	if result != 25 {
		t.Error("Expected 25", result)
	}//end if
}//end of function

func TestSubtract(t *testing.T) {
	var result int
	result = calc.Subtract(15,10)
	if result != 5 {
		t.Error("Expected 5", result)
	}
}//end of function
