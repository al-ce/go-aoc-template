package day01

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	avl "github.com/al-ce/go-avltree"

	h "go-aoc-template/internal/helpers"
)

func makeTrees(lines []string) (*avl.AvlTree[int], *avl.AvlTree[int]) {
	col1, col2 := avl.NewAvlTree[int](), avl.NewAvlTree[int]()

	for _, l := range lines {
		row := strings.Split(l, "   ")
		leftVal, err := strconv.Atoi(row[0])
		h.Check(err)
		rightVal, err := strconv.Atoi(row[1])
		h.Check(err)
		col1.Add(leftVal)
		col2.Add(rightVal)
	}
	return col1, col2
}

func PartTwo(lines []string) string {
	col1, col2 := makeTrees(lines)
	iter1, iter2 := col1.NewIterator(), col2.NewIterator()

	similarity := map[int]int{}
	var val int
	var sum int
	for range col2.Size() {
		val, _ = iter2.Next()
		_, exists := similarity[val]
		if !exists {
			similarity[val] = 0
		}
		similarity[val] += 1
	}

	for range col1.Size() {
		val, _ = iter1.Next()
		sum += val * similarity[val]
	}

	return fmt.Sprintf("%d", sum)
}

func PartOne(lines []string) string {
	col1, col2 := makeTrees(lines)
	iter1, iter2 := col1.NewIterator(), col2.NewIterator()
	var sum int
	for range col1.Size() {
		left, _ := iter1.Next()
		right, _ := iter2.Next()
		sum += int(math.Abs(float64(left) - float64(right)))
	}

	return fmt.Sprintf("%d", sum)
}
