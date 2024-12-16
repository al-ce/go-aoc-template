package day11

import (
	"fmt"
	"math"

	h "go-aoc-template/internal/helpers"
)

type Split struct {
	left  int
	right int
}

var splits = map[int]Split{}

func splitStone(stone, digits int) (int, int) {
	if split, exists := splits[stone]; exists {
		return split.left, split.right
	}
	mid := digits / 2
	left, right := 0, 0
	for i := range mid {
		right += (stone % 10) * int(math.Pow10(i))
		stone /= 10
		left = stone
	}
	return left, right
}


func blinkStone(stone, blinks int, cache map[string]int) int {
	if blinks == 0 {
		return 1
	}
	key := fmt.Sprintf("%d %d", stone, blinks)
	if count, exists := cache[key]; exists {
		return count
	} else {
		cache[key] = 0
	}
	digits := len(fmt.Sprintf("%d", stone))
	count := 0
	if stone == 0 {
		count += blinkStone(1, blinks-1, cache)
	} else if digits%2 == 0 {
		left, right := splitStone(stone, digits)
		count += blinkStone(left, blinks-1, cache)
		count += blinkStone(right, blinks-1, cache)
	} else {
		count += blinkStone(stone*2024, blinks-1, cache)
	}
	cache[key] = count
	return count
}

func solver(lines []string, blinks int) string {
	cache := make(map[string]int)
	count := 0
	stones := h.ParseIntString(lines[0], " ")
	for _, stone := range stones {
		count += blinkStone(stone, blinks, cache)
	}
	return fmt.Sprintf("%d", count)
}

func PartOne(lines []string) string {
	return solver(lines, 25)
}

func PartTwo(lines []string) string {
	return solver(lines, 75)
}
