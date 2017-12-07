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
