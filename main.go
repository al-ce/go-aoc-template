package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	h "go-aoc-template/internal/helpers"
	solutions "go-aoc-template/internal/imports"
)

func parser(args []string) (int, int) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Need exactly two args: <day> <part>\n")
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

	part, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not parse %s\n", args[0])
		os.Exit(1)
	}
	if !(part == 1 || part == 2) {
		fmt.Fprintf(os.Stderr, "%s: part must be 1 or 2\n", args[0])
		os.Exit(1)
	}

	return day, part
}

func main() {
	flag.Parse()
	args := flag.Args()
	day, part := parser(args)
	input := solutions.InputCoords{Day: day, Part: part}

	solution, exists := solutions.Solutions[input]
	if !exists {
		fmt.Fprintf(os.Stderr, "No solution implemented for day %d\n", day)
		os.Exit(1)
	}

	lines, err := h.InputLines(day)
	h.Check(err)
	fmt.Println(solution(lines))
}
