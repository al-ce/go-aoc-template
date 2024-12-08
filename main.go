package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	h "go-aoc-template/internal/helpers"
	solutions "go-aoc-template/internal/imports"
)

func parser(args []string) int {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <day>\n", os.Args[0])
		os.Exit(1)
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not parse %s\n", args[0])
		os.Exit(1)
	}
	if day < 1 || day > 25 {
		fmt.Fprintf(os.Stderr, "%s: day must be 1-25\n", args[0])
		os.Exit(1)
	}

	return day
}

func runSolution(day int, part int, lines []string) {
	input := solutions.InputCoords{Day: day, Part: part}
	solutionFunc, exists := solutions.Solutions[input]
	if !exists {
		fmt.Fprintf(os.Stderr, "No solution implemented for day %d\n", day)
		os.Exit(1)
	}
	timelog_start_1 := time.Now()
	solution := solutionFunc(lines)
	elapsed := time.Since(timelog_start_1)
	fmt.Printf("Part %d: %s (%v)\n", part, solution, elapsed)
}

func main() {
	flag.Parse()
	args := flag.Args()
	day := parser(args)
	lines, err := h.InputLines(day)
	h.Check(err)

	runSolution(day, 1, lines)
	runSolution(day, 2, lines)
}
