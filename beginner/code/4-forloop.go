package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}

	fmt.Println("\nWithin for loop ...")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}

	j := 0
	fmt.Println("\nWithin infinite for loop ...")
	for {
		if j > len(arr) {
			break
		}

		fmt.Println(j)
		j = j + 1
	}
}
