package arrays

import "fmt"

// creating a fixed-length array and basic indexing.
func Intro() {
	fmt.Println("== Arrays – Introduction ==")
	var nums [5]int
	nums[0] = 10
	fmt.Println(nums) // [10 0 0 0 0]
	fmt.Println()
}
