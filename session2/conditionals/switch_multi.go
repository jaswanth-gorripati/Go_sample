package conditionals

import "fmt"

// switch with multiple cases mapping to same branch.
func SwitchMulti() {
	fmt.Println("== Switch with Multiple Cases ==")
	grade := "B"
	switch grade {
	case "A", "B":
		fmt.Println("Excellent")
	case "C":
		fmt.Println("Average")
	default:
		fmt.Println("Needs Improvement")
	}
	fmt.Println()
}
