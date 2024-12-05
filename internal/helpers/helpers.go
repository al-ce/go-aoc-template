package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

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
