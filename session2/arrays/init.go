package arrays

import "fmt"

// array literal initialization.
func Init() {
	fmt.Println("== Arrays – Initialization ==")
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
	fmt.Println()
}
