package conditionals

import "fmt"

// switch without an expression (acts like chained if).
func SwitchNoExpr() {
	fmt.Println("== Switch Without Expression ==")
	num := -5
	switch {
	case num < 0:
		fmt.Println("Negative")
	case num == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive")
	}
	fmt.Println()
}
