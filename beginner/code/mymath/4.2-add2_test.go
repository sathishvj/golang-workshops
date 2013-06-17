package mymath

import (
	"testing"
)

func Test_Add2(t *testing.T) {
	s := Add2(5, 10)
	if s != 15 {
		t.Errorf("FAIL: Expected 15.  Received %d.", s)
	} else {
		t.Logf("PASS: Expected 15, also received %d.", s)
	}

	s = Add2(15, 10)
	if s != 25 {
		t.Errorf("FAIL: Expected 25.  Received %d.", s)
	} else {
		t.Logf("PASS: Expected 25, also received %d.", s)
	}
}

// 1. package testing
// 2. Import a pointer to it in every test.
// 3. The test file ends with "_test.go"
// 4. Each test begins with Test_
// 5. Use the function "t *testing.T" to get to the testing functionality
// 6. Run tests with "go test" inside the directory with the files
// 7. 
