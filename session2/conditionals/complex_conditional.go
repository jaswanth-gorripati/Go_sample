package conditionals

import "fmt"

// Complex boolean logic in a conditional.
func Complex() {
	fmt.Println("== Complex Conditional Example ==")
	age := 30
	hasTicket := true
	if age >= 18 && hasTicket {
		fmt.Println("Allowed to enter concert")
	} else {
		fmt.Println("Not allowed")
	}
	fmt.Println()
}
