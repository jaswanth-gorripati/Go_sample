package main

import (
	"database/sql"
	"day2/util"
	"fmt"

	"github.com/google/uuid" // Importing SQLite driver for database operations
)

func NewID() string {
	return uuid.New().String()
}

func main() {
	fmt.Printf("add(5, 3) = %d\n", util.Add(5, 3))
	fmt.Printf("subtract(5, 3) = %d\n", util.Subtract(5, 3))
	fmt.Printf("multiply(5, 3) = %d\n", util.Multiply(5, 3))
	fmt.Printf("divide(5, 3) = %.2f\n", util.Divide(5, 3))
	_, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

}
