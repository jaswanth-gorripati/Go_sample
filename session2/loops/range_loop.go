package loops

import "fmt"

// iterating using range.
func RangeLoop() {
	fmt.Println("== Loop with range ==")
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}
	fmt.Println()
}
