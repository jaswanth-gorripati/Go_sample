package loops

import "fmt"

// nested loops
func Nested() {
	fmt.Println("== Nested Loops ==")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()
}
