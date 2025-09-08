package main

import (
	"fmt"
)

type Student struct {
	Roll  int
	Name  string
	Age   int
	Mark  int
	marks map[string]int
}

func main() {

	studentMap := map[int]*Student{}

	stu1 := Student{Name: "Alice", Age: 20, Mark: 85}
	var stu2 Student
	stu2.Name = "Bob"
	stu2.Age = 22
	stu2.Mark = 90

	studentMap[1111] = &stu1
	studentMap[1112] = &stu2

	fmt.Println(studentMap[1112])

}
