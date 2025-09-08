package main

import (
	"flag"
	"fmt"
	"os"

	"session2/arrays"
	"session2/conditionals"
	"session2/loops"
)

func runLoops() {
	loops.Classic()
	loops.WhileStyle()
	loops.InfiniteDemo()
	loops.BreakContinue()
	loops.Nested()
	loops.RangeLoop()
	loops.SumEvens()
}

func runConditionals() {
	conditionals.IfElse()
	conditionals.ShortStmt()
	conditionals.SwitchBasic()
	conditionals.SwitchMulti()
	conditionals.SwitchNoExpr()
	conditionals.Complex()
}

func runArrays() {
	arrays.Intro()
	arrays.Init()
	arrays.Properties()
	arrays.Iterating()
	arrays.CricketScoreboard()
}

func list() {
	fmt.Println("Available demo groups:")
	fmt.Println("  loops        - classic, while-style, infinite demo, break/continue, nested, range, sum evens")
	fmt.Println("  cond         - if/else, short-stmt, switch (basic/multi/no-expr), complex condition")
	fmt.Println("  arrays       - intro, init, properties, iterating, cricket scoreboard")
}

func main() {
	demo := flag.String("demo", "", "run only a specific demo group: loops | cond | arrays")
	listFlag := flag.Bool("list", false, "list available demos and exit")
	flag.Parse()

	if *listFlag {
		list()
		return
	}

	switch *demo {
	case "loops":
		runLoops()
	case "cond":
		runConditionals()
	case "arrays":
		runArrays()
	case "":
		// default: run all safe demos
		fmt.Println("Running ALL demos. Use --demo to filter or --list to see options.")
		runLoops()
		runConditionals()
		runArrays()
	default:
		fmt.Println("Unknown demo:", *demo)
		list()
		os.Exit(1)
	}
}
