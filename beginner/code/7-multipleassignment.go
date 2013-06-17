package main

import (
	"fmt"
	"strconv"
)

func SumProd(i, j int) (int, int) {
	return i + j, i * j
}

func main() {
	s, p := SumProd(5, 6)
	fmt.Println(s, p)

	arr := []string{"Hello", "how", "are", "you?"}
	for i, v := range arr 
		fmt.Println(i, v)
	}

	a := "20a"
	if _, err := strconv.Atoi(a); err != nil {
		fmt.Println("Error! ", err)
	}

}

// 1. Multiple assignment possible
// 2. Functions must have return variables in bracket.
// 3. You can also name return variables, but not show here.  Eg.. func SumProd(i, j int) (sum, prod int) { ... }
// 4. "range" keywords gives you a (position, value at position) tuple
// 5. You can use range also with maps, which are the equivalent of hashtables
// 6. The same multiple return value is used in an error checking paradigm
