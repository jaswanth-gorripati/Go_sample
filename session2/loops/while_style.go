package loops

import "fmt"

// while-style loop example.
func WhileStyle() {
	fmt.Println("== While-Style Loop ==")
	i := 1
	for i <= 5 {
		fmt.Println("Number:", i)
		i++
	}
	fmt.Println()
}
