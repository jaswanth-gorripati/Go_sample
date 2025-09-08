package arrays

import "fmt"

// Properties zero values, length, and bounds.
func Properties() {
	fmt.Println("== Arrays – Properties & Default Values ==")
	var marks [5]int
	fmt.Println(marks) // [0 0 0 0 0]
	// accessing out-of-range like marks[5] would panic
	fmt.Println("Length:", len(marks))
	fmt.Println()
}
