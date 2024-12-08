package helpers

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Read a file and return a slice of all lines
func InputLines(day int) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("inputs/%d", day))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find input file for day %d. Add input to ./inputs/%d\n", day, day)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Take a string that contains ints separated by `sep` and return a slice of
// ints. Ex: "1 2 3" returns []int {1, 2 3}
func ParseIntString(str string, sep string) []int {
	split := strings.Split(str, sep)
	intSlice := make([]int, len(split))
	for i, v := range split {
		converted, err := strconv.Atoi(v)
		Check(err)
		intSlice[i] = converted
	}
	return intSlice
}

// Get the absolute diff between two integers
func GetAbsDiff(l, r int) int {
	return int(math.Abs(float64(l - r)))
}

// Get the direction of change between two integers
func GetSignOfDiff(l, r int) int {
	diff := GetAbsDiff(l, r)
	if diff == 0 {
		return 0
	}
	return (l - r) / -diff
}

// Check if a number is within a range, inclusive
func InRange(x, min, max int) bool {
	return x >= min && x <= max
}
