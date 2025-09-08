package loops

import (
	"fmt"
	"time"
)

// "infinite" loop but exits after 3 iterations to remain safe.
func InfiniteDemo() {
	fmt.Println("== Infinite Loop (demo, 3 iterations) ==")
	count := 0
	for {
		fmt.Println("Waiting for user input...")
		time.Sleep(200 * time.Millisecond)
		count++
		if count >= 3 {
			fmt.Println("Stopping demo to avoid an actual infinite loop.")
			break
		}
	}
	fmt.Println()
}
