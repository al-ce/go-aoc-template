package day13

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func makeCoord(s string) Coord {
	pattern := regexp.MustCompile(`(\d+),\s\w.(\d+)`)
	matches := pattern.FindStringSubmatch(s)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return Coord{x, y}
}

func getPrize(buttonA, buttonB, prize Coord, memo map[Coord]int) int {
	if prize.x < 0 || prize.y < 0 {
		return -1
	} else if prize.x == 0 && prize.y == 0 {
		return 0
	}
	if solved, exists := memo[prize]; exists {
		return solved
	}

	minTokens := math.MaxInt

	nextPrizeA := Coord{prize.x - buttonA.x, prize.y - buttonA.y}
	countA := getPrize(buttonA, buttonB, nextPrizeA, memo)
	if countA >= 0 && countA < minTokens {
		minTokens = countA + 3
	}

	nextPrizeB := Coord{prize.x - buttonB.x, prize.y - buttonB.y}
	countB := getPrize(buttonA, buttonB, nextPrizeB, memo)
	if countB >= 0 && countB < minTokens {
		minTokens = countB + 1
	}

	if minTokens == math.MaxInt {
		memo[prize] = -1
	} else {
		memo[prize] = minTokens
	}
	return memo[prize]
}

func PartOne(lines []string) string {
	machines := strings.Split(strings.Join(lines, "\n"), "\n\n")
	total := 0
	for _, m := range machines {
		_m := strings.Split(m, "\n")
		a, b, p := _m[0], _m[1], _m[2]
		buttonA := makeCoord(a)
		buttonB := makeCoord(b)
		prize := makeCoord(p)
		memo := make(map[Coord]int)
		thisMachine := getPrize(buttonA, buttonB, prize, memo)
		if thisMachine != -1 {
			total += thisMachine
		}
	}
	return fmt.Sprintf("%d", total)
}

func PartTwo(lines []string) string {
	return "TODO"
}
