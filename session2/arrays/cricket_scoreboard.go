package arrays

import "fmt"

// A 2D array representing overs x balls.
func CricketScoreboard() {
	fmt.Println("== Complex Array Example - Cricket Scoreboard ==")
	scores := [3][6]int{
		{1, 0, 4, 6, 1, 2}, // Over 1
		{0, 2, 1, 6, 4, 0}, // Over 2
		{1, 1, 2, 0, 6, 4}, // Over 3
	}
	for over := 0; over < 3; over++ {
		fmt.Printf("Over %d: ", over+1)
		total := 0
		for ball := 0; ball < 6; ball++ {
			run := scores[over][ball]
			fmt.Printf("%d ", run)
			total += run
		}
		fmt.Printf(" | Total in over: %d\n", total)
	}
	fmt.Println()
}
