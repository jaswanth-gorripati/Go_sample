package loops

import "fmt"

// continue and break.
func BreakContinue() {
	fmt.Println("== Break & Continue ==")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			// skip 5
			continue
		}
		if i == 8 {
			// stop at 8
			break
		}
		fmt.Println(i)
	}
	fmt.Println()
}
