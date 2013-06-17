package main

import "fmt"

var gi int

func main() {
	fmt.Println(gi) //0

	var i int
	fmt.Println(i) //0
	i = 25
	fmt.Println(i) //25

	j := 5
	s := "Hello!"
	fmt.Println("The two values are:", j, s)
	//The two values are: 5, Hello

	fmt.Printf("The integer is %d, and the string is %s.\n", j, s)
	//The integer is 5, and the string is Hello.

	var arr1 []int
	arr1 = []int{1, 2, 3, 4}
	arr2 := []int{1, 2, 3, 4}
	fmt.Println(arr1, arr2) //[1,2,3,4] [1,2,3,4]
}

// 1. Inverse of other ways of declaration.  C/Java - int i = 10;
// 2. Default values / zero values
// 3. := definition and initialization
// 4. Println automatically formats data type for printing
// 5. Formatted prints with Printf
// 6. array of values
