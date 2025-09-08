package main

import (
	"fmt"
)

type User struct {
	Name string
}

func counter() func(int) int {
	i := 0
	return func(x int) int {
		i += x
		return i
	}
}

func Operate(x, y int, op func(int, int) int) int {
	result := op(x, y)
	return result
}
func main() {
	defer fmt.Println("Done")

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(Operate(3, 4, add))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(Operate(3, 4, sub))
	defer fmt.Println("Exiting")
	counter := counter()
	for i := 0; i < 5; i++ {
		fmt.Println(counter(i))
	}
	defer fmt.Println("Cleaning up")

	//var done bool
	// var str string
	// var i int
	// if done == "user Given value" && done == false {
	// 	fmt.Println("Not done yet")
	// } else {
	// 	fmt.Println("All done")
	// }

	// s := []int{1, 2, 3, 4, 5}
	// s = s[:len(s)-1]

	// p := new(string)
	// *p = 42
	// fmt.Println(*p)
	// u := User{Name: "Alice"}

	// r := make(map[string]int, 5)
	// r["one"] = 1
	// fm

}
