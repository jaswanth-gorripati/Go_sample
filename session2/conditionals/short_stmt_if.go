package conditionals

import "fmt"

// if with a short variable declaration.
func ShortStmt() {
	fmt.Println("== Short Statement in If ==")
	if num := 15; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	fmt.Println()
}
