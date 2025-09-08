package conditionals

import "fmt"

// if / else if / else grading example.
func IfElse() {
	fmt.Println("== If / Else If / Else ==")
	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 70 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Fail")
	}
	fmt.Println()
}
