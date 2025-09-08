package conditionals

import "fmt"

// simple switch on a day.
func SwitchBasic() {
	fmt.Println("== Switch Statement ==")
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("Work week starts")
	case "Friday":
		fmt.Println("Weekend coming")
	default:
		fmt.Println("Mid-week day")
	}
	fmt.Println()
}
