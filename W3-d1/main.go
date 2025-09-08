package main

import "fmt"

func printAll[T any](v T) T {
	fmt.Println(v)
	return v
}

func main() {
	fmt.Println(printAll("kjdfkj"))
	printAll(0)

}
