package arrays

import "fmt"

// Iterating both index-based and range-based iteration over arrays.
func Iterating() {
	fmt.Println("== Iterating Over Arrays ==")
	marks := [5]int{90, 85, 70, 88, 95}

	fmt.Println("-- for i := 0; i < len(marks); i++ --")
	for i := 0; i < len(marks); i++ {
		fmt.Println("Student", i+1, "Scored:", marks[i])
	}

	fmt.Println("-- for i, val := range marks --")
	for i, val := range marks {
		fmt.Println("Index:", i, "Value:", val)
	}
	fmt.Println()
}
