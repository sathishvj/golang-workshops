package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// 1. package statement starts of all .go files
// 2. "main" should be the package name for an executable file.  Library files use other names, e.g. "package math"
// 3. You include/import other packages with the import statement
// 4. Functions are defined with the "func" keyword
// 5. An executable file starts executing at the "main" function.  Note that both the package name and the starting function are both called "main"
// 6. Library functions are used with the dot operator
// 7. TRY: what happens if you put fmt.println("Hello World")
// 8. Check documentation at golang.org/pkg/fmt

// 9. Conventions in coding with go
// 10.
