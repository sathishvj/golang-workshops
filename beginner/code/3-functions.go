package main

import (
	"fmt"
)

func Add(i, j int) int {
	return i + j
}

func main() {
	s := Add(5, 10)
	fmt.Println("Sum is: ", s)
}

// 1. function takes parameters in a similar fashion - variable name followed by type
// 2. "return" statement
