package loops

import "fmt"

// Classic for loop example.
func Classic() {
	fmt.Println("== Classic For Loop ==")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	fmt.Println()
}
