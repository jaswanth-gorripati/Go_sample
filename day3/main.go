package main

import (
	"fmt"
	"runtime"
	"time"
)

type snapshot struct {
	heapAlloc uint64
	mallocs   uint64
}

func memSnap() snapshot {
	runtime.GC() // force a GC so the numbers are cleaner/comparable
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return snapshot{heapAlloc: m.HeapAlloc, mallocs: m.Mallocs}
}

func human(b uint64) string {
	const kb = 1024
	const mb = 1024 * 1024
	if b >= mb {
		return fmt.Sprintf("%.2f MB", float64(b)/float64(mb))
	}
	if b >= kb {
		return fmt.Sprintf("%.2f KB", float64(b)/float64(kb))
	}
	return fmt.Sprintf("%d B", b)
}

func runCase(title string, initialCap, inserts, step int) {
	fmt.Printf("\n=== %s (initialCap=%d) ===\n", title, initialCap)
	m := make(map[int]int, initialCap)
	base := memSnap()
	last := base

	start := time.Now()
	for i := 0; i < inserts; i++ {
		m[i] = i
		if (i+1)%step == 0 {
			now := memSnap()
			fmt.Printf("len=%-8d  HeapAlloc=%-10s  ΔHeap=%-9s  Mallocs=%-10d  ΔMallocs=%-8d\n",
				len(m),
				human(now.heapAlloc),
				human(now.heapAlloc-last.heapAlloc),
				now.mallocs,
				now.mallocs-last.mallocs,
			)
			last = now
		}
	}
	total := memSnap()
	elapsed := time.Since(start)

	fmt.Printf("---- SUMMARY ----\n")
	fmt.Printf("Final len: %d\n", len(m))
	fmt.Printf("Total HeapAlloc: %s (Δ from start: %s)\n", human(total.heapAlloc), human(total.heapAlloc-base.heapAlloc))
	fmt.Printf("Total Mallocs: %d (Δ from start: %d)\n", total.mallocs, total.mallocs-base.mallocs)
	fmt.Printf("Elapsed: %v\n", elapsed)
}

func main() {
	const inserts = 200000
	const step = 1000

	runCase("No capacity hint", 0, inserts, step)
	runCase("With capacity hint", inserts, inserts, step)
}
