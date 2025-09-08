package loops

import "fmt"

// sum of even numbers between 1 and 100.
func SumEvens() {
	fmt.Println("== Sum of even numbers 1..100 ==")
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of evens:", sum)
	fmt.Println()
}
