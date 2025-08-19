package main

import "fmt"

// Lower case function name -- Not exported
// Upper case function name -- Exported
// func <function_name> (parameters type, ....) (return_type, ... ) {
// 	// Function body
// 	fmt.Println("This is a function in main package")
// }

func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to the Go programming world.\n", name)
}
