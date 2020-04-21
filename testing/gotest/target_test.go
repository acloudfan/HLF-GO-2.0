package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {

	// 1. Execute the target  test function
	result := Calculate(2)

	// 2. Check the result
	if result != 4 {

		// 3. Indicate Test Case Failure with a log message
		t.Error("Expected 2 + 2 to equal 4")
	}

	// 4. Following msg will be logged too
	t.Logf("This is a %s that will appear only if there is an error or -v is used", "message")
}
