package day13

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x float64
	y float64
}

func makeCoord(s string) Coord {
	pattern := regexp.MustCompile(`(\d+),\s\w.(\d+)`)
	matches := pattern.FindStringSubmatch(s)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return Coord{float64(x), float64(y)}
}

func solver(lines []string, part int) string {
	machines := strings.Split(strings.Join(lines, "\n"), "\n\n")

	total := 0
	for _, m := range machines {

		_m := strings.Split(m, "\n")

		a, b, p := _m[0], _m[1], _m[2]
		buttonA := makeCoord(a)
		buttonB := makeCoord(b)
		prize := makeCoord(p)
		if part == 2 {
			prize.x += 10000000000000
			prize.y += 10000000000000
		}

		// solve using row echelon form
		m := [][]float64{
			{buttonA.x, buttonB.x, prize.x},
			{buttonA.y, buttonB.y, prize.y},
		}

		// Set 1,0 to 1 by dividing row 1 by m[1][0]
		temp := []float64{
			1, m[1][1] / m[1][0], m[1][2] / m[1][0],
		}
		// Interchange rows
		m = [][]float64{temp, m[0]}
		// Set 0, 0 to 0 by adding row0 * (1, 0) to row 1
		m[1] = []float64{
			0,
			m[0][1]*-m[1][0] + m[1][1],
			m[0][2]*-m[1][0] + m[1][2],
		}
		// Set 1, 1 to 1 by dividing row 1 by m[1][1]
		m[1] = []float64{
			0, 1, m[1][2] / m[1][1],
		}

		// Solve for button variables
		B := math.Round(m[1][2])
		A := math.Round(m[0][2] - (m[0][1] * B))

		if buttonA.x*A+buttonB.x*B == prize.x &&
			buttonA.y*A+buttonB.y*B == prize.y {
			cost := A*3 + B
			total += int(cost)
		}
	}

	return fmt.Sprintf("%d", int(total))
}

func PartOne(lines []string) string {
	return solver(lines, 1)
}

func PartTwo(lines []string) string {
	return solver(lines, 2)
}
