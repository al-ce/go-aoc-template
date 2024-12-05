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

func InputLines(day, part int) ([]string, error) {
	file, err := os.Open(fmt.Sprintf("inputs/day%d-%d", day, part))
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
